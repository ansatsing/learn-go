package condition

import "testing"

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i { //类似java语言
		case 0, 2:
			t.Log("Odd:", i)
		case 1, 3:
			t.Log("Even:", i)
		default:
			t.Log("Unkown:", i)
		}
	}

	for i := 0; i < 5; i++ {
		switch { //类似 if else if
		case i%2 == 0:
			t.Log("Odd:", i)
		case i%2 == 1:
			t.Log("Even:", i)
		}
	}
}
