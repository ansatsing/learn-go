package ch9_func

import (
	"fmt"
	"testing"
)

//defer:delay excuted function,类似java的finally

func clear() {
	fmt.Println("clear resources!")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("start exucting data insertion.....")
}

func TestDeferWithPanic(t *testing.T) {
	defer clear()
	fmt.Println("start exucting data insertion.....")
	panic("err") //抛出异常，但clear函数还是会执行。
}
