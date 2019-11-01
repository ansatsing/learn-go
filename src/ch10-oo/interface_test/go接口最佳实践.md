### Go接口最佳实践
- 倾向于使用最小接口定义，很多接口只包含一个方法
```
type Reader interface {
    Read(p []byte) (n int, err error)
 }

type Writer interface {
    Write(p []byte) (n int, err error)
}
```
- 较大接口的定义，可以由多个小接口定义组合而成
```
type ReadWriter interface {
    Reader
    Writer
}
```
- 只依赖于必要功能的最小接口
```
func StoreData(reader Reader) error {
     …
}
```
