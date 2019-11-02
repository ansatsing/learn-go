package ch14_Groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i) //每次打印的顺序都不一样
		}(i)
		time.Sleep(time.Second * 2) //每2秒
	}
}
