package type_test

import "testing"



//类型转换是非常严格


func TestNotSupportImplicitCovert(t *testing.T){
	var a int32 = 1
	var b int64
	//b = a //会报编译错误：cannot use a (type int32) as type int64 in assignment,说明不支持这样的转换，java语言是支持这样的

	b = int64(a)

	t.Log(b)
}

type MyInt int64
func TestNotSupportSameTypeAliasCovert(t *testing.T){
	var b int64 = 2
	var a MyInt
	//a = b //编译报错：cannot use b (type int64) as type MyInt in assignment
	a = MyInt(b)

	t.Log(a)
}
