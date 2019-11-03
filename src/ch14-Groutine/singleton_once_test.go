package ch14_Groutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type MyObj struct {
}

var myObj *MyObj
var myObj1 MyObj
var once sync.Once

func getSingletonMyObj() *MyObj {
	once.Do(func() {
		myObj = new(MyObj)
	})
	return myObj
}
func getSingletonMyObj1() MyObj {
	once.Do(func() {
		myObj1 = MyObj{}
	})
	return myObj1
}

func TestOnce(t *testing.T) {
	var myObj2 MyObj
	myObj2 = getSingletonMyObj1()
	for i := 0; i < 5; i++ {
		go func() {
			//fmt.Printf("%x\n",unsafe.Pointer(getSingletonMyObj()))
			fmt.Printf("%x\n", unsafe.Pointer(&myObj2)) //不能这样&getSingletonMyObj1()
		}()
	}
	time.Sleep(time.Second * 1)
}
