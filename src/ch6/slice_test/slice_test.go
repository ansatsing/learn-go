package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var slc0 = []int{}
	t.Log(slc0, len(slc0), cap(slc0)) // [] 0 0

	slc0 = append(slc0, 2)
	t.Log(slc0, len(slc0), cap(slc0)) // [2] 1 1

	var slc1 = []int{1, 2, 3, 4}
	t.Log(slc1, len(slc1), cap(slc1)) // [1 2 3 4] 4 4

	slc2 := make([]int, 4, 9)
	t.Log(slc2, len(slc2), cap(slc2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(s, len(s), cap(s))
	}
}

//有待深入 todo
func TestSliceShareMemory(t *testing.T) {
	month := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	q2 := month[3:6]
	summer := month[5:8]
	t.Log(q2, len(q2), cap(q2))
	t.Log(summer, len(summer), cap(summer))

	summer[0] = 14
	t.Log(q2, len(q2), cap(q2))
	t.Log(summer, len(summer), cap(summer))
	t.Log(month, len(month), cap(month))

	summer[3] = 18
	t.Log(q2, len(q2), cap(q2))
	t.Log(summer, len(summer), cap(summer))
	t.Log(month, len(month), cap(month))

}
