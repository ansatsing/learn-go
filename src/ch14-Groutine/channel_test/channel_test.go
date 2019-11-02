package channel_test

import (
	"fmt"
	"testing"
	"time"
)

//运行报错：fatal error: all goroutines are asleep - deadlock!
func TestChannel(t *testing.T) {
	chn := make(chan int) //非缓存通道
	chn <- 10
	fmt.Println(<-chn)
}

func TestChannel1(t *testing.T) {
	chn := make(chan int, 1) //缓存通道
	chn <- 10
	fmt.Println(<-chn)
}

//fatal error: all goroutines are asleep - deadlock!
func TestChanne2(t *testing.T) {
	chn := make(chan int) //非缓存通道
	chn <- 10             //这个位置阻塞了
	fmt.Println("..........")
	go func() {
		fmt.Println(<-chn)
	}()
}

func TestChanne3(t *testing.T) {
	chn := make(chan int) //非缓存通道
	go func() {
		chn <- 10
	}()

	fmt.Println(<-chn)
}

func TestChannel4(t *testing.T) {
	chn := make(chan string)
	fmt.Println(<-chn) //fatal error: all goroutines are asleep - deadlock!，因为一直阻塞到通道有个值可读为止
	fmt.Println("...........")
}

//fatal error: all goroutines are asleep - deadlock!
func TestChannel5(t *testing.T) {
	chn := make(chan string)
	chn <- "jjjjj" //fatal error: all goroutines are asleep - deadlock!,一直阻塞直到被读走
}

func TestChannel6(t *testing.T) {
	chn := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println(<-chn) //执行完这行代码，，，立马执行60行代码
	}()
	chn <- "jjjjj" //fatal error: all goroutines are asleep - deadlock!,一直阻塞直到被读走
	fmt.Println(".......")
}
