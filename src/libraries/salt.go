package libraries

import (
	"crypto/rand"
	"fmt"
	"time"
)

// 產生 Salt 字串
func MakeSalt() string {

	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 4)
	num, err := rand.Read(buff)
	if num != len(buff) || err != nil {
		panic(err)
	}
	salt := fmt.Sprintf("%x-%x-%x", unix32bits, buff[0:2], buff[2:4])

	return salt
}