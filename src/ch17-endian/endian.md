# big endian符合人的思维习惯：低地址存高位，高地址存地位，通常都是大端
## 例子
```
十进制： 305419896
16进制: 0x12345678
2进制: 00010010-00110100-01010110-01111000  高位--->低位
```
## litle endian
<html>
       低地址 ------------------> 高地址
       #################################################### 
值    #0111，1000 #0101，0110 # 0011，0100  #  0001，0010  #
       ####################################################
地址 #      100        #     101        #          102       #      103          # 
       ####################################################
</html>

## big endian
<html>
       低地址 ------------------> 高地址
       #################################################### 
 值   # 0001，0010  #0011，0100 # 0101，0110  #  0111，1000  #
       ####################################################
地址 #       100         #     101         #        102        #        103         # 
       ####################################################
</html>

# 另一篇讲解
### 字节序概述
> 主机字节序模式有两种，大端数据模式和小端数据模式，在网络编程中应注意这两者的区别，以保证数据处理的正确性；例如网络的数据是以大端数据模式进行交互，而我们的主机大多数以小端模式处理，如果不转换，数据会混乱 参考 ；一般来说，两个主机在网络通信需要经过如下转换过程：主机字节序 —> 网络字节序 -> 主机字节序
### 大小端区别
```
大端模式：Big-Endian就是高位字节排放在内存的低地址端，低位字节排放在内存的高地址端
低地址 --------------------> 高地址
高位字节                     地位字节
小端模式：Little-Endian就是低位字节排放在内存的低地址端，高位字节排放在内存的高地址端
低地址 --------------------> 高地址
低位字节                     高位字节
```
### 什么是高位字节和低位字节
```
例如在32位系统中，357转换成二级制为：00000000 00000000 00000001 01100101，其中

00000001 | 01100101 
高位字节     低位字节
```
### int和byte转换
```
在go语言中，byte其实是uint8的别名，byte 和 uint8 之间可以直接进行互转。目前来只能将0~255范围的int转成byte。因为超出这个范围，go在转换的时候，就会把多出来数据扔掉;如果需要将int32转成byte类型，我们只需要一个长度为4的[]byte数组就可以了
```
### 大端模式
```
func f2() {
    var v2 uint32
    var b2 [4]byte
    v2 = 257
    // 将 257转成二进制就是
    // | 00000000 | 00000000 | 00000001 | 00000001 |
    // | b2[0]    | b2[1]    | b2[2]    | b2[3]    | // 这里表示b2数组每个下标里面存放的值
    // 这里直接使用将uint32强转成uint8
    // | 00000000 0000000 00000001 | 00000001  直接转成uint8后等于 1
    // |---这部分go在强转的时候扔掉---|
    b2[3] = uint8(v2)
    // | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
    // 下面是右移后的数据
    // |          | 00000000 | 00000000 | 00000001 |
    b2[2] = uint8(v2 >> 8)
    // | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
    // 下面是右移后的数据
    // |          |          | 00000000 | 00000000 |
    b2[1] = uint8(v2 >> 16)
    // | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
    // 下面是右移后的数据
    // |          |          |          | 00000000 |
    b2[0] = uint8(v2 >> 24)
    fmt.Printf("%+v\n", b2)
    // 所以最终将uint32转成[]byte数组输出为
    // [0 0 1 1]
}
```
### 小端模式下
```// 在上面我们讲过，小端刚好和大端相反的，所以在转成小端模式的时候，只要将[]byte数组的下标首尾对换一下位置就可以了
func f3() {
  var v3 uint32
  var b3 [4]byte
  v3 = 257
  // 将 256转成二进制就是
  // | 00000000 | 00000000 | 00000001 | 00000001 |
  // | b3[0]  | b3[1]  | b3[2]  | [3]   | // 这里表示b3数组每个下标里面存放的值
  // 这里直接使用将uint32l强转成uint8
  // | 00000000 0000000 00000001 | 00000001 直接转成uint8后等于 1
  // |---这部分go在强转的时候扔掉---|
  b3[0] = uint8(v3)
  // | 00000000 | 00000000 | 00000001 | 00000001 | 右移8位 转成uint8后等于 1
  // 下面是右移后的数据
  // |     | 00000000 | 00000000 | 00000001 |
  b3[1] = uint8(v3 >> 8)
  // | 00000000 | 00000000 | 00000001 | 00000001 | 右移16位 转成uint8后等于 0
  // 下面是右移后的数据
  // |     |     | 00000000 | 00000000 |
  b3[2] = uint8(v3 >> 16)
  // | 00000000 | 00000000 | 00000001 | 00000001 | 右移24位 转成uint8后等于 0
  // 下面是右移后的数据
  // |     |     |     | 00000000 |
  b3[3] = uint8(v3 >> 24)
  fmt.Printf("%+v\n", b3)
  // 所以最终将uint32转成[]byte数组输出为
  // [1 1 0 0 ]
}
```
### go转换demo
```
//整形转换成字节
func IntToBytes(n int) []byte {
  x := int32(n)
  bytesBuffer := bytes.NewBuffer([]byte{})
  binary.Write(bytesBuffer, binary.BigEndian, x)
  return bytesBuffer.Bytes()
}
//字节转换成整形
func BytesToInt(b []byte) int {
  bytesBuffer := bytes.NewBuffer(b)
 
  var x int32
  binary.Read(bytesBuffer, binary.BigEndian, &x)
 
  return int(x)
}
```