package exchange

import "testing"

func TestExchangeValue(t *testing.T) {
	a := 1
	b := 2
	t.Log("before exchange:")
	t.Log(a, b)  //1 2
	a, b = b, a
	t.Log("after exchange:")
	t.Log(a, b)	//2 1
}
