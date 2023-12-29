package models

import (
	"fmt"
	"testing"
	"unsafe"
)

// Go内存对齐

type s1 struct {
	a int8
	b string
	c int8
}

type s2 struct {
	a int8
	c int8
	b string
}

func TestStruct(t *testing.T) {
	v1 := s1{
		a: 1,
		b: "Eutop1a",
		c: 2,
	}

	v2 := s2{
		a: 1,
		b: "Eutop1a",
		c: 2,
	}

	fmt.Println(unsafe.Sizeof(v1)) // 32
	fmt.Println(unsafe.Sizeof(v2)) // 24
}
