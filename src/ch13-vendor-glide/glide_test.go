package ch13_vendor_glide

import "testing"
import cm "github.com/easierway/concurrent_map"

//如果没有没有执行命令：glide install，则编译时会报错：
//FAIL	ch13-vendor-glide [setup failed]
//# ch13-vendor-glide
//package ch13-vendor-glide (test)
//imports github.com/easierway/concurrent_map: cannot find package "github.com/easierway/concurrent_map" in any of:
//C:\Go\src\github.com\easierway\concurrent_map (from $GOROOT)
//E:\go_workspace\src\github.com\easierway\concurrent_map (from $GOPATH)
//E:\GolandProjects\learn-go\src\github.com\easierway\concurrent_map
//FAIL
func TestGlide(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"), 10)
	t.Log(m.Get(cm.StrKey("key")))
}
