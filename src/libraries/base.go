package libraries

import "encoding/base64"

func Base64Encrypt(data string) string {
    return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64Decrypt(baseCode string) (string, error) {
    result, err := base64.URLEncoding.DecodeString(baseCode)
    return string(result), err
}