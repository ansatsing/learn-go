package pipeline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

func ArraySource(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			//time.Sleep(time.Millisecond*200)
			out <- v
		}
		close(out) //一般都不需要，但养成close的习惯
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		// read into memory
		a := []int{}
		for v := range in {
			a = append(a, v)
		}

		//sort
		sort.Ints(a)

		//output
		for _, v := range a {
			out <- v
		}
		//time.Sleep(time.Millisecond*200)
		// close(out)
	}()
	return out
}

func Merge(chn1, chn2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-chn1
		v2, ok2 := <-chn2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-chn1
			} else {
				out <- v2
				v2, ok2 = <-chn2
			}
		}
		close(out) //all goroutines are asleep - deadlock!
	}()
	return out
}

//读取数据
func ReaderSource(reader io.Reader) <-chan int {
	out := make(chan int)
	go func() {
		buf := make([]byte, 8)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buf))
				out <- v
			}
		}
		close(out)
	}()
	return out
}

//持久化数据
func WriterSink(writer io.Writer, in <-chan int) {
	buf := make([]byte, 8)
	for v := range in {
		binary.BigEndian.PutUint64(buf, uint64(v))
		writer.Write(buf)
	}
}

//随机数据源
func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
