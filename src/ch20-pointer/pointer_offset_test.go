package main

import (
	"fmt"
	"testing"
	"unsafe"
)

//测试通过uinpointer加偏移量来修改   结构体字段值

type Man1 struct {
	xb     bool
	salary int32
	name   string
}

func TestPointerOffset(t *testing.T) {
	//man:=Man1{true,89,"syq"}
	//p:=unsafe.Pointer(&man)
	//namePtr:=(*string)(unsafe.Pointer(uintptr(p)+unsafe.Offsetof(man.name)))
	//fmt.Println(man.name)
	//*namePtr="nba"
	//fmt.Println(man.name)

	/////////////////
	man1 := Man1{true, 89, "syq"}
	var manPtr *Man1
	manPtr = &man1
	var manIntPtr uintptr
	manIntPtr = uintptr(unsafe.Pointer(manPtr))

	var manNamePtr uintptr
	manNamePtr = manIntPtr + unsafe.Offsetof(manPtr.name)
	fmt.Println(man1.name)
	//*manNamePtr = "nba" //invalid indirect of manNamePtr (type uintptr)
	manNameIntPtr := unsafe.Pointer(manNamePtr)
	//*manNameIntPtr="nba" //invalid indirect of manNameIntPtr (type unsafe.Pointer)
	manNameStringPtr := (*string)(manNameIntPtr) //因为name属性是字符串，所以一定把uintptr转换成字符串类型指针,否则没发通过指针来修改name的值
	*manNameStringPtr = "nba"
	fmt.Println(man1.name)

}
