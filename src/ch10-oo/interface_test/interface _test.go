package interface_test

import (
	"fmt"
	"testing"
)

type Pgrogrammer interface {
	helloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) helloWorld() string {
	return "fmt.println(\"hello world\")"
}

type GoProgrammer1 struct {
}

func (g *GoProgrammer1) helloWorld() string {
	return "fmt1.println(\"hello world\")"
}

func TestInterface(t *testing.T) {
	var programmer Pgrogrammer
	programmer = new(GoProgrammer)
	t.Log(programmer.helloWorld())
}

func TestCheckType(t *testing.T) {
	var programmer Pgrogrammer
	programmer = new(GoProgrammer1)
	Check(programmer)
	//var i = 0
	//switch  o:=i.(type){  // 报错
	//
	//}
}

func Check(p Pgrogrammer) {
	switch r := p.(type) { //针对接口
	case *GoProgrammer:
		r.helloWorld()
		fmt.Println("GoProgrammer ", r)
	case *GoProgrammer1:
		fmt.Println("GoProgrammer1 ", r)

	default:
		fmt.Println("unknow ", r)
	}
}
