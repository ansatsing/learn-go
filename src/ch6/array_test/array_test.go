package array_test

import (
	"reflect"
	"testing"
)

func TestArray(t *testing.T) {
	var a = [...]int{1, 2, 3, 4}                //数组
	t.Log(a, len(a), cap(a), reflect.TypeOf(a)) //[1 2 3 4] 4 4

	var b = []int{1, 2, 3, 4}                   //称 切片
	t.Log(b, len(b), cap(b), reflect.TypeOf(b)) //[1 2 3 4] 4 4

	var c = [...]int{1, 2, 3, 4}
	if a == c {
		t.Log("a == c：", a == c)
	}

	//if a == b { //编译报错：invalid operation: a == b (mismatched types [4]int and []int)

	//}
}
