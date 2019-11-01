package libraries

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)
// 因為snowFlake目的是解決分散式下生成唯一id 所以ID中是包含叢集和節點編號在內的
const (
	numberBits uint8 = 12 // 表示每個叢集下的每個節點，1毫秒內可生成的id序號的二進位制位 對應上圖中的最後一段
	workerBits uint8 = 8  // 每臺機器(節點)的ID位數 10位最大可以有2^10=1024個節點數 即每毫秒可生成 2^12-1=4096個唯一ID 對應上圖中的倒數第二段
	// 這裡求最大值使用了位運算，-1 的二進位制表示為 1 的補碼，感興趣的同學可以自己算算試試 -1 ^ (-1 << nodeBits) 這裡是不是等於 1023
	workerMax   int64 = -1 ^ (-1 << workerBits) // 節點ID的最大值，用於防止溢位
	numberMax   int64 = -1 ^ (-1 << numberBits) // 同上，用來表示生成id序號的最大值
	timeShift   uint8 = workerBits + numberBits // 時間戳向左的偏移量
	workerShift uint8 = numberBits              // 節點ID向左的偏移量
	// 41位位元組作為時間戳數值的話，大約68年就會用完
	// 假如你2010年1月1日開始開發系統 如果不減去2010年1月1日的時間戳 那麼白白浪費40年的時間戳啊！
	// 這個一旦定義且開始生成ID後千萬不要改了 不然可能會生成相同的ID
	epoch int64 = 1566202889901 // 這個是我在寫epoch這個常量時的時間戳(毫秒)
)

type UUID int64

var timestamp int64
var number int64
var workerId int64
var mutex sync.Mutex

// 生成方法一定要掛載在某個woker下，這樣邏輯會比較清晰 指定某個節點生成id
func GetSnowId() UUID {

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	ipL := strings.LastIndex(localAddr, ".")

	workerId, _ = strconv.ParseInt(localAddr[ipL+1:idx], 10, 64)

	// 獲取id最關鍵的一點 加鎖 加鎖 加鎖
	mutex.Lock()
	defer mutex.Unlock() // 生成完成後記得 解鎖 解鎖 解鎖

	// 獲取生成時的時間戳
	now := time.Now().UnixNano() / 1e6 // 納秒轉毫秒
	if timestamp == now {
		number++

		// 這裡要判斷，當前工作節點是否在1毫秒內已經生成numberMax個ID
		if number > numberMax {
			number = 0
			// 如果當前工作節點在1毫秒內生成的ID已經超過上限 需要等待1毫秒再繼續生成
			for now <= timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		// 如果當前時間與工作節點上一次生成ID的時間不一致 則需要重置工作節點生成ID的序號
		number = 0
		// 下面這段程式碼看到很多前輩都寫在if外面，無論節點上次生成id的時間戳與當前時間是否相同 都重新賦值  這樣會增加一丟丟的額外開銷 所以我這裡是選擇放在else裡面
		timestamp = now // 將機器上一次生成ID的時間更新為當前時間
	}

	uid := int64((now-epoch)<<timeShift | (workerId << workerShift) | (number))

	return UUID(uid)
}

func (u UUID) String() string {
	idstr := strconv.FormatInt(int64(u), 10)
	return idstr
}

func (u UUID) Hexstring() string {
	idstr := strconv.FormatInt(int64(u), 16)
	idstr = "0000000000000000" + idstr
	idstr = idstr[len(idstr)-16 : len(idstr)]
	return idstr
}

func UuidFromHexString(uuidString string) int64 {

	uuidInt64, _ := strconv.ParseInt(uuidString, 16, 64)

	return uuidInt64
}

func TimeToShift(time int64) int64 {

	shiftTime := int64(time << timeShift)

	return shiftTime
}
