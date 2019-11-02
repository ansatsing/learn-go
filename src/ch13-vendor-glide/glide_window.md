## 环境
```
E:\GolandProjects\learn-go>ver

Microsoft Windows [版本 10.0.17134.523]

E:\GolandProjects\learn-go>go version
go version go1.13.4 windows/amd64

```
## 安装
- go get -u github.com/Masterminds/glide
- 查看首个工作区bin目录
```
E:\go_workspace\bin>dir
 驱动器 E 中的卷没有标签。
 卷的序列号是 76D7-3EF9

 E:\go_workspace\bin 的目录

2019/11/02  11:04    <DIR>          .
2019/11/02  11:04    <DIR>          ..
2019/11/02  11:04        12,490,240 glide.exe
2019/10/29  23:09         4,864,000 goaddimport.exe
2019/10/29  23:09         7,432,704 gogetdoc.exe
2019/10/29  23:11         4,527,104 gomodifytags.exe
2019/10/29  23:11         2,203,136 gopkgs.exe
2019/10/29  22:05         1,956,864 hello.exe
               6 个文件     33,474,048 字节
               2 个目录 112,118,480,896 可用字节


```
## 使用
### 1 glide init
```
E:\GolandProjects\learn-go\src\ch13-vendor-glide>glide init
[INFO]  Generating a YAML configuration file and guessing the dependencies
[INFO]  Attempting to import from other package managers (use --skip-import to skip)
[INFO]  Scanning code to look for dependencies
[INFO]  --> Found test reference to github.com\easierway\concurrent_map
[INFO]  Writing configuration file (glide.yaml)
[INFO]  Would you like Glide to help you find ways to improve your glide.yaml configuration?
[INFO]  If you want to revisit this step you can use the config-wizard command at any time.
[INFO]  Yes (Y) or No (N)?
Y
[INFO]  Looking for dependencies to make suggestions on
[INFO]  --> Scanning for dependencies not using version ranges
[INFO]  --> Scanning for dependencies using commit ids
[INFO]  Gathering information on each dependency
[INFO]  --> This may take a moment. Especially on a codebase with many dependencies
[INFO]  --> Gathering release information for dependencies
[INFO]  --> Looking for dependency imports where versions are commit ids
[INFO]  Here are some suggestions...
[INFO]  The package github.com/easierway/concurrent_map appears to have Semantic Version releases (http://semver.org).
[INFO]  The latest release is 0.9.1. You are currently not using a release. Would you like
[INFO]  to use this release? Yes (Y) or No (N)
Y
[INFO]  Would you like to remember the previous decision and apply it to future
[INFO]  dependencies? Yes (Y) or No (N)
N
[INFO]  Updating github.com/easierway/concurrent_map to use the release 0.9.1 instead of no release
[INFO]  The package github.com/easierway/concurrent_map appears to use semantic versions (http://semver.org).
[INFO]  Would you like to track the latest minor or patch releases (major.minor.patch)?
[INFO]  The choices are:
[INFO]   - Tracking minor version releases would use '>= 0.9.1, < 1.0.0' ('^0.9.1')
[INFO]   - Tracking patch version releases would use '>= 0.9.1, < 0.10.0' ('~0.9.1')
[INFO]   - Skip using ranges
[INFO]  For more information on Glide versions and ranges see https://glide.sh/docs/versions
[INFO]  Minor (M), Patch (P), or Skip Ranges (S)?
S
[INFO]  Would you like to remember the previous decision and apply it to future
[INFO]  dependencies? Yes (Y) or No (N)
N
[INFO]  Configuration changes have been made. Would you like to write these
[INFO]  changes to your configuration file? Yes (Y) or No (N)
y
[INFO]  Writing updates to configuration file (glide.yaml)
[INFO]  You can now edit the glide.yaml file.:
[INFO]  --> For more information on versions and ranges see https://glide.sh/docs/versions/
[INFO]  --> For details on additional metadata see https://glide.sh/docs/glide.yaml/

E:\GolandProjects\learn-go\src\ch13-vendor-glide>
```
### 2查看glide.yaml
```
E:\GolandProjects\learn-go\src\ch13-vendor-glide>type glide.yaml
package: ch13-vendor-glide
import: []
testImport:
- package: github.com/easierway/concurrent_map
  version: 0.9.1

```
### 3glide install
```
E:\GolandProjects\learn-go\src\ch13-vendor-glide>glide install
[INFO]  Lock file (glide.lock) does not exist. Performing update.
[INFO]  Downloading dependencies. Please wait...
[INFO]  --> Fetching github.com/easierway/concurrent_map
[INFO]  --> Setting version for github.com/easierway/concurrent_map to 0.9.1.
[INFO]  Resolving imports
[INFO]  Downloading dependencies. Please wait...
[INFO]  Setting references for remaining imports
[INFO]  Exporting resolved dependencies...
[INFO]  --> Exporting github.com/easierway/concurrent_map
[INFO]  Replacing existing vendor dependencies
[ERROR] Unable to export dependencies to vendor directory: Error moving files: exit status 1. output: �ܾ����ʡ�
�ƶ��ˡ�        0 ��Ŀ¼��

```
> 报错 Unable to export dependencies to vendor directory: Error moving files: exit status 1. output
>。解决方法：解决办法： 找到%GOPATH%/src/github.com/Masterminds/glide/path/winbug.go 文件，修改约75行处
```
func CustomRename(o, n string) error {

// Handking windows cases first
if runtime.GOOS == "windows" {
    msg.Debug("Detected Windows. Moving files using windows command")
    //cmd := exec.Command("cmd.exe", "/c", "move", o, n)
    cmd := exec.Command("cmd.exe", "/c", "xcopy /s/y", o, n+"\\") //新增这一行代码
    output, err := cmd.CombinedOutput()
    if err != nil {
        return fmt.Errorf("Error moving files: %s. output: %s", err, output)
    }

    return nil
} else if detectWsl() {
    cmd := exec.Command("mv", o, n)
    output, err2 := cmd.CombinedOutput()
    msg.Debug("Detected Windows Subsystem for Linux. Removing files using subsystem command")
    if err2 != nil {
        return fmt.Errorf("Error moving files: %s. output: %s", err2, output)
    }

    return nil
}

return os.Rename(o, n)
}
```
> 重新编译后[go build glide.go]复制glide.exe文件到%GOPATH%/bin目录下【重要】,再次glide install就OK了
```
E:\GolandProjects\learn-go\src\ch13-vendor-glide>glide install
[INFO]  Lock file (glide.lock) does not exist. Performing update.
[INFO]  Downloading dependencies. Please wait...
[INFO]  --> Fetching updates for github.com/easierway/concurrent_map
[INFO]  --> Setting version for github.com/easierway/concurrent_map to 0.9.1.
[INFO]  Resolving imports
[INFO]  Downloading dependencies. Please wait...
[INFO]  Setting references for remaining imports
[INFO]  Exporting resolved dependencies...
[INFO]  --> Exporting github.com/easierway/concurrent_map
[INFO]  Replacing existing vendor dependencies
[INFO]  Project relies on 0 dependencies.

```
### 4查看目录结构
```
E:\GolandProjects\learn-go\src\ch13-vendor-glide>tree /f
文件夹 PATH 列表
卷序列号为 76D7-3EF9
E:.
│  glide.lock
│  glide.yaml
│  glide_test.go
│  glide_window.md
│  readme.md
│
└─vendor
    └─github.com
        └─easierway
            └─concurrent_map
                    concurrent_map.go
                    concurrent_map_benchmark_adapter.go
                    concurrent_map_test.go
                    int64_key.go
                    LICENSE
                    map_benchmark.png
                    map_benchmark_test.go
                    README.md
                    rwlock_map_benchmark_adapter.go
                    string_key.go
                    sync_map_benchmark_adapter.go

```
### 5再测试glide_test.go就不会报编译错误了

## 参考资料
- https://github.com/Masterminds/glide
- http://www.mamicode.com/info-detail-2271636.html
