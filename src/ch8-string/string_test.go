package ch8_string

import (
	"testing"
	"unsafe"
)

func TestStringIsImmutableByteSlice(t *testing.T) {
	var a string
	t.Log(a, len(a)) //  0
	a = "hello"
	t.Log(a, len(a)) // hello 5
	//a[1]=1  //编译错误： cannot assign to a[1]，不可更改，string是不可变的byte slice

	a = "\xE4\xB8\xA5" //可以存储任何二进制数据
	t.Log(a, len(a))   // 严 3
}

//			字符串编码与存储
//		字符				中
//		Unicode				0x4E2D
//		UTF-8				0xE4B8AD
//	  string/[]byte			[0xE4,0xB8,0xAD]
func TestStringUnicodeUtf8(t *testing.T) {
	//var s ='中'  //s的类型是 rune
	//t.Logf("s len:%d",len(s)) // invalid argument s (type rune) for len

	var s string
	s = "中"
	t.Logf("s len:%d", len(s)) //s len:3   字节长度为3

	u := []rune(s)                              //获取中字的unicode  u[1]会报index out of range的错误
	t.Log("rune u size: ", unsafe.Sizeof(u[0])) //("'中'的unicode %x",u[0])
	t.Logf("'中'的unicode: %x", u[0])             //'中'的unicode: 4e2d
	t.Logf("'中'的UTF8: %x", s)                   //'中'的UTF8: e4b8ad

}
