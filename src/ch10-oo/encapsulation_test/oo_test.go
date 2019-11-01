package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

//类似java class，但只包含属性，不包含行为
type Employee struct {
	id   string
	name string
	age  int
}

//每执行一次，就需要拷贝一份，如果出于数据安全考虑，推荐此写法
func (e Employee) String() string {
	return fmt.Sprintf("Employee[id:%s,name:%s,age:%d],address:%x", e.id, e.name, e.age, unsafe.Pointer(&e))
}

//无需拷贝 ，都是同一份，如果不更改数据，推荐此写法
func (e *Employee) StringOfPointer() string {
	return fmt.Sprintf("Employee[id:%s,name:%s,age:%d],address:%x", e.id, e.name, e.age, unsafe.Pointer(e))
}

func TestEmployee(t *testing.T) {
	e1 := Employee{"001", "antsing", 35}
	e2 := Employee{id: "002", name: "syq", age: 33}
	e3 := new(Employee) //返回的是指针  类似&Employee{"001","antsing",35}
	e3.age = 32
	e3.name = "sjq"
	e3.id = "003"
	t.Log("e1: ", e1) //{001 antsing 35}
	t.Log("e2: ", e2)
	t.Log("e3: ", e3) // &{003 sjq 32}

	t.Logf("type of e2: %T", e2) //type of e2: ch10_oo.Employee
	t.Logf("type of e3: %T", e3) //type of e3: *ch10_oo.Employee
}

func TestEmployeeBehavior(t *testing.T) {
	e := Employee{"001", "antsing", 35}
	t.Logf("Employee[id:%s,name:%s,age:%d],address:%x", e.id, e.name, e.age, unsafe.Pointer(&e))
	t.Log(e.String())
	t.Log(e.StringOfPointer())
}
