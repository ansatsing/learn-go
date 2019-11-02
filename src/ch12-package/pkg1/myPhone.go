package pkg1

import "fmt"

func init() {
	fmt.Println("myPhone init0")
}

func init() {
	fmt.Println("myPhone init1")
}

//方法名大写开头意味着包外的代码可以访问
func Sing() {
	fmt.Println("myPhone is singing......")
}

func close() {
	fmt.Println("myPhone is closing....")
}
