package util

import (
	"encoding/base64"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	origData := []byte("{\"test_key\", \"test_body\"}")
	key := []byte("test0000000000000000000000000000")
	iv := []byte("abcdefghijklmnop")

	res, err := AesEncrypt(origData, key, iv)
	if err != nil {
		t.Error(err)
	}

	if base64.StdEncoding.EncodeToString(res) != "PLgBVh+X6mGF4pzToGpN2qnqrTrpb2Xltv3LBUwrtq4=" {
		t.Error("加密失败")
	}
}
