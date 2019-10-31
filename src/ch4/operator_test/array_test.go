package operator_test

import "testing"

func TestArraySupportSameDimensionAndLenCompare(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c:=[...]int{1,2,3,0,3}
	d := [...]int{1, 2, 3, 4}
	t.Log(d, len(d), cap(d))

	t.Log(a == b)
	//t.Log(a > b) //不支持这样的比较
	//t.Log(a == c) //编译报错：nvalid operation: a == c (mismatched types [4]int and [5]int)，维度相同，但长度不一样
	t.Log(a == d)

	f := a[1:3]
	t.Log(f, len(f), cap(f))

	e := [5]int{1, 2, 3, 4, 5}
	t.Log(e, len(e), cap(e))
}
