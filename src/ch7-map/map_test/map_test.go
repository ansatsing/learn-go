package map_test

import "testing"

func TestMapInit(t *testing.T) {
	var m1 = map[string]int{"a": 1, "b": 2, "c": 3, "'\"d\"'": 4}
	//t.Log(m1,len(m1),cap(m1)) //编译报错：invalid argument m1 (type map[string]int) for cap
	t.Log(m1, len(m1)) //map[a:1 b:2 c:3] 3

	m1["f"] = 5
	t.Log(m1, len(m1))

	m2 := make(map[int]int, 2)
	t.Log(m2, len(m2))
	for i := 0; i < 5; i++ {
		m2[i] = i
	}
	t.Log(m2, len(m2))
}

func TestMapTravel(t *testing.T) {
	var m1 = map[string]int{"a": 1, "b": 2, "c": 3, "'\"d\"'": 4}
	for key, value := range m1 {
		t.Log(key, value)
	}
}

func TestKeyExisting(t *testing.T) {
	var m1 = map[int]int{}
	t.Log(m1[0]) // 0
	t.Log(m1[3]) // 0
	if v, exist := m1[2]; exist {
		t.Log("key[2] is existing,value:", v)
	} else {
		t.Log("key[2] is not existing")
	}
	m1[3] = 4
	if v, exist := m1[3]; exist {
		t.Log("key[3] is existing,value:", v)
	} else {
		t.Log("key[3] is not existing")
	}
	var m2 = map[int]string{}
	t.Log(m2[0]) //""
	t.Log(m2[3]) //""
}
