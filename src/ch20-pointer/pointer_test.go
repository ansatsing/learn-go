package main

import (
	"fmt"
	"testing"
)

type Man struct {
	flag   bool
	age    int16
	salary float32
}

func TestPointerTypeSecurityCheck(t *testing.T) {
	man := new(Man)
	man.flag = false
	man.age = 23
	man.salary = 23455.33

	//var pointer *int32
	//pointer = *int32(&man.age)  //int16类型的指针不能转换成int32的指针类型，出于安全考虑，否则会乱套，知道内存地址就可以随意更改值

	var p1 *int16
	p1 = &man.age
	*p1 = 33
	fmt.Println(man.age)

	type MyInt int32
	var myInt MyInt = 33
	var i int32
	//i=myInt   //cannot use myInt (type MyInt) as type int32 in assignment
	i = int32(myInt)
	fmt.Println(i)
}
