package ch14_Groutine

import (
	"fmt"
	"testing"
	"time"
)

////////////生产者与消费者////////////
func Producer(queue chan<- int) { //chan<- 写操作
	for i := 0; i < 10; i++ {
		queue <- i
		fmt.Println("produce a data:", i)
	}
}

func Consumer(queue <-chan int) { //<-chan 读操作
	for i := 0; i < 10; i++ {
		v := <-queue
		fmt.Println("consume a data:", v)
	}
}

func TestProCon(t *testing.T) {
	queue := make(chan int, 1) //创建缓存量为1的通道
	go Producer(queue)
	go Consumer(queue)

	time.Sleep(1e9) //让Producer与Consumer完成
}

/////////////////////////////

/////////利用CSP实现类似java的future特性//////////

//模拟耗时任务
func getData(i int) int {
	time.Sleep(time.Second * 1)
	fmt.Println("挖矿结束。。。。。")
	return i
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 600)
	fmt.Println("Task is done.")
}

//总耗时--- PASS: TestCommon (1.60s)
func TestUseCommon(t *testing.T) {
	data := getData(8)
	otherTask()
	fmt.Println("data:", data)
}

func getDataByCsp(i int) chan int {
	chn := make(chan int, 1)
	go func() {
		result := getData(i)
		chn <- result
	}()
	return chn
}

//PASS: TestUseCsp (1.00s)
func TestUseCsp(t *testing.T) {
	chn := getDataByCsp(8)
	otherTask()
	fmt.Println("data:", <-chn)
	//time.Sleep(time.Second * 1)
}

//////////////////
