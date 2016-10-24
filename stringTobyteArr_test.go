package conver

import (
	"testing"
)

func Benchmark_StringToByteArray_A(b *testing.B) {
	src := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		dest := []byte(src)
		dest = dest
	}
}
func Benchmark_StringToByteArray_B(b *testing.B) {
	src := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		dest := String2ByteArr(src)
		dest = dest
	}
}

func Benchmark_ByteArrayToString_A(b *testing.B) {
	src := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < b.N; i++ {
		dest := string(src)
		dest = dest
	}
}
func Benchmark_ByteArrayToString_B(b *testing.B) {
	src := []byte("abcdefghijklmnopqrstuvwxyz")
	for i := 0; i < b.N; i++ {
		dest := ByteArr2String(src)
		dest = dest
	}
}
