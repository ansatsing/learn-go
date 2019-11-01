## Go 的错误机制
- 没有异常机制
- error 类型实现了了 error 接⼝口
```
type error interface {
Error() string
}
```
- 可以通过 errors.New 来快速创建错误实例例
```
errors.New("n must be in the range [0,100]")
```


## 最佳实践：
> 尽早失败，避免嵌套