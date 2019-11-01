package libraries

import (
    "crypto/des"
    "encoding/hex"
    "errors"
)

func DesEncrypt(src, key string, paddingFunc func([]byte, int) []byte) (string, error) {
    block, err := des.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    bs := block.BlockSize()
    src = string(paddingFunc([]byte(src),bs))
    if len(src)%bs != 0 {
        return "", errors.New("Need a multiple of the blocksize")
    }
    out := make([]byte, len(src))
    dst := out
    for len(src) > 0 {
        block.Encrypt(dst, []byte(src)[:bs])
        src = src[bs:]
        dst = dst[bs:]
    }
    return hex.EncodeToString(out), nil
}


func DesDecrypt(src, key string, unPaddingFunc func([]byte) []byte) (string, error) {
    b, _ := hex.DecodeString(src)
    src = string(b)
    block, err := des.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }
    out := make([]byte, len(src))
    dst := out
    bs := block.BlockSize()
    if len(src)%bs != 0 {
        return "", errors.New("crypto/cipher: input not full blocks")
    }
    for len(src) > 0 {
        block.Decrypt(dst, []byte(src)[:bs])
        src = src[bs:]
        dst = dst[bs:]
    }

    out = unPaddingFunc(out)
    return string(out), nil
}