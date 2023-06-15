package util

import (
	"math/rand"
	"strings"
)

var (
	letters = []byte("abcdefhkmnrstuvwxz0123456789")
	longLetters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
    base32Letters = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567")
)

// RandLowStr 随机字符串，包含 0~9 和 a~z - [g,i,j,l,o,p,q,y]
func RandLowStr(length int) string {
	if length <= 0 {
		return ""
	}

	b := make([]byte, length)

	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}

	for i := range b {
		b[i] = letters[uint8(rand.Int63()%28)]
	}

	return string(b)
}

// RandUpStr 随机字符串，包含 英文字母和数字
func RandUpStr(length int) string {
	if length <= 0 {
		return ""
	}

	b := make([]byte, length)

	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}

	for i := range b {
		b[i] = longLetters[uint8(rand.Int63()%61)]
	}

	return string(b)
}

// RandBase32Str 生成随机base32编码字符串
func RandBase32Str(length int) string {
	if length <= 0 {
		return ""
	}

	b := make([]byte, length)

	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}

	for i := range b {
		b[i] = base32Letters[uint8(rand.Int63()%28)]
	}

	return string(b)
}

// SpliceStr 拼接字符串
func SpliceStr(p ...string) string {
	var b strings.Builder
	l := len(p)
	for i := 0; i < l; i++ {
		b.WriteString(p[i])
	}
	return b.String()
}

// IsContainStr 判断字符串切片中是否存在某个字符串
func IsContainStr(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}

	return false
}
