package conver

import (
	"unsafe"
)

// go 基础类型 slice 定义
type Slice struct {
	ptr unsafe.Pointer
	len int
	cap int
}

func String2ByteArr(s string) []byte {
	slice := (*Slice)(unsafe.Pointer(&s))
	// 转化后cap 容量异常 手动修复
	slice.cap = slice.len + 1
	return *(*[]byte)(unsafe.Pointer(slice))
}

// go 基础类型 string 定义
/*
type String struct {
	ptr unsafe.Pointer
	len int
}
*/

func ByteArr2String(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}
