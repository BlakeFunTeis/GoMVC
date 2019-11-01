package libraries

import (
	"crypto/rand"
	"fmt"
	"time"
)

func MakeUUID() string {

	unix32bits := uint32(time.Now().UTC().Unix())
	buff := make([]byte, 12)
	num, err := rand.Read(buff)
	if num != len(buff) || err != nil {
		panic(err)
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x-%x", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])

	return uuid
}