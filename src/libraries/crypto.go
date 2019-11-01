package libraries

import (
    "crypto/md5"
    "encoding/hex"
    "golang.org/x/crypto/bcrypt"
)

func Md5Encrypt(data string) string {
    temp := []byte(data)
    hash := md5.New()
    hash.Write(temp)
    result := hex.EncodeToString(hash.Sum(nil))
    return result
}

func BcryptEncrypt(data string) string {
    hash, _ := bcrypt.GenerateFromPassword([]byte(data), 5)
    return string(hash)
}