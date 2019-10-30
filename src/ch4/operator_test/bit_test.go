package operator_test

import "testing"

const (
	readable = 1 << iota
	writable
	executable
)

//按位置零运算符
//	&^
//    与0运算的结果都是左边本身，与1运算的结果都是0
//   	1&^0=1
//   	0&^0=0
//   	1&^1=0
//   	0&^1=0
func TestBitClear(t *testing.T) {
	a := 7
	t.Log("是否可读:", a&readable == readable)
	t.Log("是否可写:", a&writable == writable)
	t.Log("是否可执行:", a&executable == executable)

	a = a &^ executable //清除可写权限
	t.Log("清除可写权限后在去判断")
	t.Log("是否可读:", a&readable == readable)
	t.Log("是否可写:", a&writable == writable)
	t.Log("是否可执行:", a&executable == executable)
}
