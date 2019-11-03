package ch14_Groutine

import (
	"context"
	"fmt"
	"testing"
	"time"
)

//n个任务取消方式：通过发送n消息到通道

func isCancled(chn chan struct{}) bool {
	select {
	case <-chn:
		//fmt.Println("isCancelled:",true)
		return true
	default:
		//fmt.Println("isCancelled:",false)
		return false

	}
}

func doMultipleTasks(n int, chn chan struct{}) {
	for i := 0; i < n; i++ {
		go func(j int) {
			for {
				if isCancled(chn) { //代表收到一个消息了
					fmt.Println("task", j, " is cancled!")
					break
				} else {
					//为什么2个任务都是0？？？？？？？
					//task 4  is cancled!
					//task 2  is cancled!
					//task 0  is cancled!
					//task 1  is cancled!
					//task 0  is cancled!
					time.Sleep(time.Millisecond * 100) ////如果没有这行代码就会出现多个任务j值一样
				}
			}
		}(i)

		//为什么2个任务都是0？？？？？？？
		//task 4  is cancled!
		//task 2  is cancled!
		//task 0  is cancled!
		//task 1  is cancled!
		//task 0  is cancled!
		//time.Sleep(time.Millisecond*500) //如果没有这行代码就会出现多个任务j值一样.其实不是这样的，加了这行代码经过测试还是存在同样的问题

	}
}

//这样只能有一个人被取消，因为只生产一个数据
func TestCancleAllTask(t *testing.T) {
	chn := make(chan struct{})
	doMultipleTasks(5, chn)
	time.Sleep(time.Second * 2)
	chn <- struct{}{}
}

//生产任务个数个消息，这样就全部任务被取消
//为什么2个任务都是0？？？？？？？
//task 4  is cancled!
//task 2  is cancled!
//task 0  is cancled!
//task 1  is cancled!
//task 0  is cancled!
func TestCancleAllTask1(t *testing.T) {
	chn := make(chan struct{})
	i := 5
	doMultipleTasks(i, chn)
	time.Sleep(time.Second * 2)
	for j := 0; j < i; j++ {
		chn <- struct{}{}
	}
}

//////////////////////////////////////////////////////////////////////////////////////

//n个任务取消方式：通过通道关闭的方式

func TestCancleByChannelClose(t *testing.T) {
	chn := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(j int, chn chan struct{}) {
			for {
				if isCancled(chn) { //代表通道关闭了
					fmt.Println("task ", j, " cancled")
					break
				} else {
					time.Sleep(time.Millisecond * 300)
				}
			}
		}(i, chn)
	}
	close(chn)
	time.Sleep(time.Millisecond * 900)

}

////////////////////////////////////

//n个任务取消方式：通过context

func isCancleByContext(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancledByContext(t *testing.T) {
	ctx, cancle := context.WithCancel(context.Background())

	for i := 0; i < 5; i++ {
		go func(j int, ctx context.Context) {
			for {
				if isCancleByContext(ctx) { //代表通道关闭了
					fmt.Println("task ", j, " cancled")
					break
				} else {
					time.Sleep(time.Millisecond * 300)
				}
			}
		}(i, ctx)
	}

	cancle()

	time.Sleep(time.Millisecond * 800)
}
