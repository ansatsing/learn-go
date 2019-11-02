## Thead vs. Groutine对比
#### 创建时默认的 stack的大小
- JDK5 以后 Java Thread stack 默认为1M
- Groutine 的 Stack 初始化⼤大⼩小为2K
#### 和 KSE （Kernel Space Entity) 的对应关系
- Java Thread 是 1:1
- Groutine 是 M:N