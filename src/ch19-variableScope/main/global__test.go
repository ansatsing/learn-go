package main

import (
	"fmt"
	"testing"
)

type MyInt int
type Myint1 = int

func TestGlobal(t *testing.T) {

	c := 3

	var a = MyInt(3)
	fmt.Println(a)
	//a=c  //cannot use c (type int) as type MyInt in assignment
	//fmt.Println(a)

	var b Myint1

	b = c
	fmt.Println(b)
}
