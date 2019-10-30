package type_test

import "testing"

func TestStringDefaultValueIsEmpty(t *testing.T){
	var a string
	t.Log("=="+a+"--")

	if a == "" {
		t.Log("a是空字符串")
	}
}
