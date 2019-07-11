package utils

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Encrypt(password, deskey string) (reuslt string, err error) {
	origData := []byte(password)
	key := []byte(deskey)
	block, err := des.NewCipher(key)
	if err != nil {
		return "error", err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)

	reuslt = base64.StdEncoding.EncodeToString(crypted) //一般64位编码处理
	reuslt = fmt.Sprintf("%X", crypted)                 //ODS的处理，16进制编码
	return reuslt, nil
}

func Decrypt(password, deskey string) ([]byte, error) {
	//针对ODS的反解处理，16进制字符串编码变为二进制
	crypted, _ := hex.DecodeString(password)
	// crypted := []byte(password)

	key := []byte(deskey)
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
