package remote_package

import (
	"fmt"
	"testing"
)

//github.com/ansatsing/go_repository里的内容都是src目录下的代码，但包括src本身目录
import math1 "github.com/ansatsing/go_repository/myMath" //go get -u github.com/ansatsing/go_repository/myMath，此命令就是把远程包导入本地工作区

func TestRemotePackage(t *testing.T) {
	fmt.Println(math1.Sum(1, 2))
}
