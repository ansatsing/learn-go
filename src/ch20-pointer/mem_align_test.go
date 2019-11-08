package main

import (
	"fmt"
	"testing"
	"unsafe"
)

//内存对齐
//32位操作系统 4个字节对齐
//64位操作系统 8个字节对齐

type P struct {
	flag bool //占一个字节
}

//奇怪 没有出现对齐情况，地址都是连续的
func TestMemoryAlign(t *testing.T) {
	p1 := &P{false}
	p2 := &P{true}
	p3 := &P{true}

	fmt.Printf("p1 size[%d],addr[%v]\n", unsafe.Sizeof(p1), &p1.flag)
	fmt.Printf("p2 size[%d],addr[%v]\n", unsafe.Sizeof(p2), &p2.flag)
	fmt.Printf("p3 size[%d],addr[%v]\n", unsafe.Sizeof(p3), &p3.flag)
}
