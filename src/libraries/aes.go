package libraries

import (
    "crypto/aes"
    "crypto/cipher"
)

func AesEncrypt(origData, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockSize := block.BlockSize()
    origData = paddingFunc(origData, blockSize)
    blockMode := cipher.NewCBCEncrypter(block, iv)
    crypted := make([]byte, len(origData))
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    blockMode := cipher.NewCBCDecrypter(block, iv)
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = unPaddingFunc(origData)
    return origData, nil
}