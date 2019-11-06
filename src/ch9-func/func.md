# 函数func
- 一等公民，就是跟基本类型数据一样可以作为变量，是函数式编程的重要特征
# 高阶函数
- 函数作为另一个函数的入参
- 函数作为函数的返回结果
> 满足上面其中一个就可以称为高阶函数，高阶函数也是函数式编程的重要特征
# 函数闭包
```
package main

import (
	"errors"
	"fmt"
)

type operate func(x, y int) int

type calculateFunc func(x int, y int) (int, error)

func genCalculator(op operate) calculateFunc {
	return func(x int, y int) (int, error) {//闭包函数
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}

func main() {
	x, y := 12, 23
	op := func(x, y int) int {
		return x + y
	}

	add := genCalculator(op)
	result, err := add(x, y)
	fmt.Printf("The result: %d (error: %v)\n",
		result, err)
}
```
> genCalculator函数里的匿名函数就是闭包函数。实现闭包的意义又在哪里呢？表面上看，我们只是延迟实现了一部分程序逻辑或功能而已，但实际上，我们是在动态地生成那部分程序逻辑。我们可以借此在程序运行的过程中，根据需要生成功能不同的函数，继而影响后续的程序行为。这与 GoF 设计模式中的模板方法模式有着异曲同工之妙，不是吗？