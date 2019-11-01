package libraries

import (
    "crypto/cipher"
    "crypto/des"
    "encoding/hex"
    "errors"
    "strings"
)

func TripleDesEncrypt(origData string, key, iv []byte, paddingFunc func([]byte, int) []byte) (string, error) {
    block, err := des.NewTripleDESCipher(key)
    if err != nil {
        return "", err
    }
    orig := paddingFunc([]byte(origData), block.BlockSize())
    blockMode := cipher.NewCBCEncrypter(block, iv)
    crypted := make([]byte, len(orig))
    blockMode.CryptBlocks(crypted, orig)
    return strings.ToUpper(hex.EncodeToString(crypted)), nil
}

func TripleDesDecrypt(encrypted string, key, iv []byte, unPaddingFunc func([]byte) []byte) (string, error) {
    e, err := hex.DecodeString(strings.ToLower(encrypted))
    if err != nil {
        return "", err
    }
    block, err := des.NewTripleDESCipher(key)
    if err != nil {
        return "", err
    }
    blockMode := cipher.NewCBCDecrypter(block, iv)
    origData := make([]byte, len(e))
    blockMode.CryptBlocks(origData, e)
    origData = unPaddingFunc(origData)
    if string(origData) == "unpadding error"{
        return "", errors.New("unpadding error")
    }
    return string(origData), nil
}