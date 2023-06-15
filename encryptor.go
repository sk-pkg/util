package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"time"
)

// MD5 计算指定字符串32位md5哈希
// 如果不指定字符串，则返回当前时间的md5哈希
func MD5(str string) string {
	if str == "" {
		str = strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	h := md5.New()
	h.Write([]byte(str))

	return hex.EncodeToString(h.Sum(nil))
}

// Sha1 计算指定字符串sha1散列值
func Sha1(data string) string {
	hash := sha1.New()
	hash.Write([]byte(data))

	return hex.EncodeToString(hash.Sum([]byte("")))
}

// Hmac 计算Hmac散列值
func Hmac(key, data string) string {
	hash := hmac.New(md5.New, []byte(key))
	hash.Write([]byte(data))

	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

// PKCS7Padding PKCS7填充加密
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padText...)
}

// PKCS7UnPadding PKCS7填充解密
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])

	return origData[:(length - unPadding)]
}

// AesEncrypt AES加密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesEncrypt(origData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptedData := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptedData, origData)

	return encryptedData, nil
}

// AesDecrypt AES解密
// key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256
func AesDecrypt(encryptedData, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryptedData))
	blockMode.CryptBlocks(origData, encryptedData)
	origData = PKCS7UnPadding(origData)

	return origData, nil
}
