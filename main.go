package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Test struct {
	str string
	num int64
}

func main() {
	s := "abc"
	fmt.Println("string to bytes[]", string_2_bytes(s))
	b := []byte{'a', 'b', 'c'}
	fmt.Println("bytes[] to string", bytes_2_string(b))
}

func string_2_bytes(s string) []byte {
	sp := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bp := reflect.SliceHeader{
		Data: sp.Data,
		Len:  sp.Len,
		Cap:  sp.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bp))
}

func bytes_2_string(b []byte) string {
	bp := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sp := reflect.StringHeader{
		Data: bp.Data,
		Len:  bp.Len,
	}
	return *(*string)(unsafe.Pointer(&sp))
}
