package map_test

import "testing"

//mySet:=map[int]int{}  //编译报错： expected declaration, found mySet
var mySet = map[int]bool{} //这样不会报错
func TestMapForSet(t *testing.T) {
	//var mySet = map[int]bool{}
	t.Log("key[1] is existing? ", isExisting(mySet, 1))
	mySet[2] = true
	t.Log("key[2] is existing? ", isExisting(mySet, 2))
}

func isExisting(set map[int]bool, key int) bool {
	if set == nil {
		return false
	}
	//if key == nil {
	//	return false
	//}
	if set[key] {
		return true
	}
	return false
}
