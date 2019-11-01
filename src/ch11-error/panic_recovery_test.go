package ch11_error

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func clear() {
	if err := recover(); err != nil {
		//fmt.Errorf("error message: %s",err)
		fmt.Println("recovery from ", err)
	} else {
		fmt.Println("No error msg")
	}
	fmt.Println("clear resources!")
}

func TestNormal(t *testing.T) {
	defer clear()
	fmt.Println("start...") //执行clear()
}

func TestPanic(t *testing.T) {
	defer clear()
	fmt.Println("start...")
	panic(errors.New("OOM!!")) //执行clear()
}

func TestExit(t *testing.T) {
	defer clear()
	fmt.Println("start...")
	os.Exit(-1) //不执行clear
}
