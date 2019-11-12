package main

import (
	"ch21-parallelProcessingPipeline/pipeline"
	"fmt"
)

func main() {
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
