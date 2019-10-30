package constant

import "testing"

const (
	A = iota + 2
	B
	C
)

func TestConstant(t *testing.T) {
	t.Log(A)	//2
	t.Log(B)	//3
	t.Log(C)	//4
}

const (
	OPEN = 1 << iota
	CLOSE
	PENDING
)

func TestBitLeftMoveConstant(t *testing.T) {
 	t.Log(OPEN) 	//1	0001	2的0次方
	t.Log(CLOSE)	//2	0010	2的1次方
	t.Log(PENDING)	//4	0100	2的2次方

	a:=7
	t.Log(a&OPEN == OPEN)
	t.Log(a&CLOSE == CLOSE)
	t.Log(a&PENDING == PENDING)

}
const (
	OPEN1 = 8 >> iota
	CLOSE1
	PENDING1
)

func TestBitRightMoveConstant(t *testing.T) {
	t.Log(OPEN1) 	//8	1000	2的3次方
	t.Log(CLOSE1)	//4	0100	2的2次方
	t.Log(PENDING1)	//2	0010	2的1次方
}