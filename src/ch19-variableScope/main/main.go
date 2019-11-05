package main

import "fmt"
import . "ch19-variableScope/lib"

var s = "ssss"
var str = "syq"

//container:=[]string{"zero","one","two"} //不能这样写哦
var container = []string{"zero", "one", "two"}

func main() {

	//1测试跨包变量重命名定义
	fmt.Println("s=", s)
	Print1(s)

	//2测试同包跨代码块变量重命名
	str := 2
	{
		str := "nba"
		fmt.Println("1str:", str) //nba
		print0()                  //syq
	}
	fmt.Println("2str:", str) //2
	print0()                  //syq

	///////////////////////////////////////////////////////////////////
	//跨代码块变量重名以及类型断言表达式获取真正类型
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	_, ok1 := interface{}(container).([]string)       //类型断言表达式
	_, ok2 := interface{}(container).(map[int]string) //类型断言表达式

	if !(ok1 || ok2) {
		fmt.Printf("Error: unsupported container type:%T", container)
		return
	}

	fmt.Printf("the element is %q.(container type is %T)\n", container[1], container)

	var srcInt = int16(-255)
	dstInt := int8(srcInt)

	fmt.Println(dstInt)

	//虽然直接把一个整数值转换为一个string类型的值是可行的，但值得关注的是，被转换
	//的整数值应该可以代表一个有效的 Unicode 代码点，否则转换的结果将会是"�"（仅由高亮的
	//问号组成的字符串值）。
	//字符'�'的 Unicode 代码点是U+FFFD。它是 Unicode 标准中定义的 Replacement
	//Character，专用于替换那些未知的、不被认可的以及无法展示的字符。
	s := string(-1)
	fmt.Println("s=", s)

	////////////////////////////////////////////////////////
	//别名:
	// 		[type sss=string]的别名 就名字不一样 其他什么都一样，可以==比较
	//		[type ssss string]完全不一样，不可以==比较,但可以转换，其集合是不能转换[]string与[]ssss
	type sss = string
	var s1 sss = "ccc"
	var s2 string = "ccc"
	if s1 == s2 {
		fmt.Println("s1 == s2？", true)
	}
	//集合转换
	var j1 = []sss{"1", "2"}
	var j2 []string
	j2 = []string(j1)
	fmt.Println("j2:", j2)

	type ssss string
	var s3 ssss = "ccc"
	//if s3 == s2 {//这样编译报错：invalid operation: s3 == s2 (mismatched types ssss and string)
	//
	//}
	var s4 string
	s4 = string(s3)
	fmt.Println("s4=", s4)

	//集合转换
	var j3 = []ssss{"1", "2"}
	var j4 []string
	//j2 =[]string(j3) //报编译错误.
	fmt.Println(j3, j4)
}

func print0() {
	fmt.Println("3str:", str)
}
