package map_test

import "testing"

func TestMap2Func(t *testing.T) {
	m1 := map[int]func(num int) int{}
	m1[1] = func(num int) int { return num }
	m1[2] = func(num int) int { return num * num }
	t.Log(m1[1](2), m1[2](2))
}
