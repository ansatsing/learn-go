package interface_test

import (
	"fmt"
	"testing"
)

//空接⼝口与断⾔言
//1. 空接⼝口可以表示任何类型
//2. 通过断⾔言来将空接⼝口转换为制定类型
//v, ok := p.(int) //ok=true 时为转换成功

func CheckTypeAndConvertByIf(itf interface{}) {
	if value, ok := itf.(int); ok {
		fmt.Println("Integer: ", value)
		return
	}
	if value, ok := itf.(string); ok {
		fmt.Println("string: ", value)
		return
	}
	fmt.Println("unknow! ")
}

func CheckTypeAndConvertBySwitch(itf interface{}) {
	switch value := itf.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case string:
		fmt.Println("string: ", value)
	default:
		fmt.Println("unknow! ")
	}
}

func TestEmptyInterface(t *testing.T) {
	i := 10
	s := "10"
	f := 33.3

	CheckTypeAndConvertByIf(i)     //Integer:  10
	CheckTypeAndConvertBySwitch(s) //string:  10
	CheckTypeAndConvertBySwitch(f) //unknow!
}
