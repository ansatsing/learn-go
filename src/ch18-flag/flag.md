## flag概述
> 是go语言的标准库，主要功能是接收和解析命令行参数
## demo演示
```
F:\goWorkspace\src\ch18-flag>go build main.go

F:\goWorkspace\src\ch18-flag>main.exe --help
Usage of main.exe:
  -name string
        greeting object (default "everyone")

F:\goWorkspace\src\ch18-flag>main.exe name=syq
hello!  everyone

F:\goWorkspace\src\ch18-flag>main.exe name syq
hello!  everyone

F:\goWorkspace\src\ch18-flag>main.exe --name=syq
hello!  syq

```