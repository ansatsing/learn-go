package main

import (
	"bufio"
	"ch21-parallelProcessingPipeline/pipeline"
	"fmt"
	"os"
)

func main() {
	//MergeDemo()

	//数据源写入文件
	const FILE_NAME = "small.in"
	f, err := os.Create(FILE_NAME) //创建当前工作区根目录
	defer f.Close()
	if err != nil {
		panic(err)
	}
	const SIZE = 100000000
	chn := pipeline.RandomSource(SIZE)
	writer := bufio.NewWriter(f)
	pipeline.WriterSink(writer, chn)
	writer.Flush()

	//从文件读取
	f, err = os.Open(FILE_NAME)
	if err != nil {
		panic(err)
	}
	chn = pipeline.ReaderSource(f)
	count := 0
	for v := range chn {
		fmt.Println(v)
		if count > 100 {
			break
		}
		count++
	}
}

func MergeDemo() {
	chn1 := pipeline.InMemSort(pipeline.ArraySource(3, 2, 6, 7, 4))
	chn2 := pipeline.InMemSort(pipeline.ArraySource(8, 1, 9, 5, 0))
	chn3 := pipeline.Merge(chn1, chn2)
	for v := range chn3 {
		fmt.Println(v)
	}
	//为了查看错误根源
	//v,ok:=<-chn1
	//for ok{
	//	fmt.Println(">>>>>>>>")
	//	fmt.Println(v)
	//	v,ok=<-chn1   //如果通道生产方没有关闭通道，则当所有数据消费完了此时此行代码会抛出严重错误：fatal error: all goroutines are asleep - deadlock!
	//	fmt.Println("<<<<<<<")
	//}
}
