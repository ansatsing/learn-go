package main

import (
	flag1 "ch18-flag/flag"
	"fmt"
	"os"
)

var name *string

func init() {
	name = flag1.String("name", "everyone", "greeting object")
}

func main() {
	flag1.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag1.PrintDefaults()
	}
	flag1.Parse()
	fmt.Println("hello! ", *name)

	var str = "syq"
	print1(str)
	print2(str)
	fmt.Println(str)

	strArr := []string{"aaa", "bbbb"}
	print3(strArr[:])
	print4(strArr[:])
	fmt.Println(strArr)
}

func print1(str string) {
	str = "print1"
	fmt.Println("print1:", str)
}
func print2(str string) {
	str = "print2"
	fmt.Println("print2:", str)
}
func print3(str []string) {
	str[0] = "print1"
	fmt.Println("print1:", str)
}
func print4(str []string) {
	str[0] = "print2"
	fmt.Println("print2:", str)
}
