package ch14_Groutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//感觉跟java的synchcronized里的wait和notify差不多
func TestCondtion(t *testing.T) {
	var mailbox uint8
	//mailbox=1
	var lock sync.RWMutex
	sendCond := sync.NewCond(&lock)
	receiCond := sync.NewCond(lock.RLocker())
	max := 5
	chn := make(chan struct{}, 3)
	go func(max int) { //发送
		defer func() {
			chn <- struct{}{}
		}()
		for i := 1; i < max; i++ {
			time.Sleep(500 * time.Millisecond)
			lock.Lock()
			for mailbox == 1 {
				sendCond.Wait()
			}
			fmt.Println("sendCond:", i)
			mailbox = 1
			lock.Unlock()
			receiCond.Signal()
		}
	}(max)

	go func(max int) {
		defer func() {
			chn <- struct{}{}
			//close(chn)
		}()
		for i := 1; i < max; i++ {
			time.Sleep(500 * time.Millisecond)
			lock.RLock()
			for mailbox == 0 {
				receiCond.Wait()
			}
			fmt.Println("receiCond:", i)
			mailbox = 0
			lock.RUnlock()
			sendCond.Signal()
		}
	}(max)
	<-chn
}
