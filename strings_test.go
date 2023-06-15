package util

import (
	"fmt"
	"testing"
)

func TestRandString(t *testing.T) {
	str1 := RandUpStr(23)
	fmt.Println(str1)
	if len(str1) != 23 {
		t.Error("error length strings")
	}

	str2 := RandUpStr(23)
	fmt.Println(str2)
	if str1 == str2 {
		t.Error("error strings")
	}

	str3 := RandLowStr(64)
	fmt.Println(str3)
	if len(str3) != 64 {
		t.Error("error length strings")
	}

	str4 := RandLowStr(64)
	fmt.Println(str4)
	if str3 == str4 {
		t.Error("error strings")
	}
}

func BenchmarkRandLow64String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandLowStr(64)
	}
}

func BenchmarkRandUp64String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandUpStr(64)
	}
}

func TestSpliceStr(t *testing.T) {
	a := "aaa"
	b := "bbb"
	c := "ccc"

	if "aaabbbccc" != SpliceStr(a, b, c) {
		t.Error("SpliceStr failed")
	}
}

func BenchmarkSpliceStr(b *testing.B) {
	str := "0"
	for i := 0; i < b.N; i++ {
		str = SpliceStr(str, "0")
	}
}
