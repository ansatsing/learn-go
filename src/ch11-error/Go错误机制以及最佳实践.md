## Go 的错误机制
- 没有异常机制
- error 类型实现了了 error 接⼝口
```
type error interface {
Error() string
}
```
- 可以通过 errors.New 来快速创建错误实例例
```
errors.New("n must be in the range [0,100]")
```


## 最佳实践：
> 尽早失败，避免嵌套

## 多个defer执行顺序
> FILO,先进后出，，，在defer语句每次执行的时候，Go 语言会把它携带的defer函数及其参数值
              另行存储到一个队列中。
              这个队列与该defer语句所属的函数是对应的，并且，它是先进后出（FILO）的，相当于一个
              栈。
              在需要执行某个函数中的defer函数调用的时候，Go 语言会先拿到对应的队列，然后从该队列
              中一个一个地取出defer函数及其参数值，并逐个执行调用。
              这正是我说“defer函数调用与其所属的defer语句的执行顺序完全相反”的原因了
```
package main

import "fmt"

func main() {
	defer fmt.Println("first defer")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer in for [%d]\n", i)
	}
	defer fmt.Println("last defer")
}

```