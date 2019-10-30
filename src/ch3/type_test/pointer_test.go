package type_test

import "testing"

func TestPointerNotSupportOperate(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	//t.Log("%T  %T", a, aPtr)  //这样报编译错误：Log call has possible formatting directive %T,说白了就是不支持格式化
	t.Logf("%T  %T", a, aPtr)  //类似fmt.print 和fmt.printf的区别

	//aPtr = aPtr + 1 //指针不支持运行，有别于其他语言，编译报错： invalid operation: aPtr + 1 (mismatched types *int and int)
}
