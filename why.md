## 1.空死循环必须加上time.Sleep
> 空死循环就是什么都不做，类似下面的代码
```
func doMultipleTasks(n int,chn chan struct{})  {
	for i:=0;i<n;i++{
		go func(j int) {
			for {
				if isCancled(chn) {
					fmt.Println("task",j," is cancled!")
					break
				}else {

					//为什么2个任务都是0？？？？？？？
					//task 4  is cancled!
					//task 2  is cancled!
					//task 0  is cancled!
					//task 1  is cancled!
					//task 0  is cancled!
					time.Sleep(time.Millisecond*100) ////如果没有这行代码就会出现多个任务j值一样
				}
			}
		}(i)
		//time.Sleep(time.Millisecond*500) //如果没有这行代码就会出现多个任务j值一样.其实不是这样的，加了这行代码经过测试还是存在同样的问题
	}
}
```
## 2main方法结束了其他协程也随之结束生命
> 重点关注最后一行代码
```
package main

import (
	"fmt"
	"time"
)

func isCancled(chn chan struct{}) bool{
	select{
	case <-chn:
		return true
	default:
		return false
	}
}
func main() {
	chn:=make(chan struct{})
	for i:=0;i<5;i++{
		go func(j int,cn chan struct{}) {
			for {
				if  isCancled(cn) {
					fmt.Println(j," cancled..")
					break
				}else{
					fmt.Println(j," do.....")
					time.Sleep(time.Millisecond*300)
				}
			}

		}(i,chn)
	}
	close(chn)
	//time.Sleep(time.Second*3) //不加这代码 看不到演示效果：子线程取消时的打印
}

```
## 3.&后面只能跟变量，不能跟返回变量的方法
```
详细代码参考：src/ch14-Groutine/singleton_once_test.go
func TestOnce(t *testing.T)  {
	var myObj2 MyObj
	myObj2=getSingletonMyObj1()
	for i:=0;i<5;i++{
		go func() {
			//fmt.Printf("%x\n",unsafe.Pointer(getSingletonMyObj()))
			fmt.Printf("%x\n",unsafe.Pointer(&myObj2))  //不能这样&getSingletonMyObj1()
		}()
	}
	time.Sleep(time.Second*1)
}
```

## 类型推断、变量重声明、可重名变量、类型断言、类型转换、别名类型和潜在类型