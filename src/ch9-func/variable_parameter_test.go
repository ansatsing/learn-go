package ch9_func

import "testing"

func sum(nums ...int) int {
	ret := 0
	for _, vle := range nums {
		ret += vle
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
}
