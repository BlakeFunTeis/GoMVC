# Go-MVC

### Go Modules
	$ cd Go-MVC
	$ go mod download

所有外部依賴套件通通安裝在你設定的 $GOPATH/pkg/mod。

## XORM 產生結構體
### Windows
	$ cd src
	$ xorm.exe reverse mysql root:root@(IP:PORT)/TABLE_NAME?charset=utf8 goxorm
### macOS
	$ cd src
	$ xorm reverse mysql "root:root@tcp(IP:PORT)/TABLE_NAME?charset=utf8" goxorm
goxorm => 你的$GOPATH/pkg內的github.com/go-xorm/cmd/xorm/templates/goxorm  
所有結構體會產生在src內的models


### 檔案目錄結構
    - src
        - config 設定檔
        - controllers 控制器
        - core 核心函數
        - docs API文件
        - helpers 輔助函數
        - libraries 函式庫
        - middleware 中介器
        - migrations 資料庫遷移
        - models 資料庫設定
        - repositories 資料庫邏輯
        - routes 路由器
        - schedules 排程
        - services 服務
        - template 樣板
        - tests 單元測試
	- presenter 顯示邏輯
	- notification 通知
        - go.mod 外部依賴管理文件
        - main.go 主執行程序
