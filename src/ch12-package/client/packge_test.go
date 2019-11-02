package client

import (
	math1 "ch12-package"
	妈妈 "ch12-package/pkg1"
	"fmt"
	"testing"
)

func init() {
	fmt.Println("package_test init....")
}

func TestPackage(t *testing.T) {
	妈妈.Sing() //必须把目录E:\GolandProjects\learn-go加到GOPATH里去，否则找不到依赖，window环境多个工作区用英文分号隔开
	//pkg1.close()  //会报编译错误
	fmt.Println(math1.Sum(2, 3))
}
