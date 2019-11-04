package ch17_endian

import (
	"encoding/binary"
	"fmt"
	"log"
	"testing"
	"unsafe"
)

const INT_SIZE int = int(unsafe.Sizeof(0))

//检查系统属于哪种方式存储字节序列
func TestCheckSystemEndian(t *testing.T) {
	//t.Log(INT_SIZE)  //8
	//t.Log(unsafe.Sizeof(0)) //8
	//
	//a:=[1]int64{}
	//t.Log(unsafe.Sizeof(a)) //8
	//t.Log(len(a),cap(a)) // 1 1
	//
	//b:=int64(4)
	//t.Log(unsafe.Sizeof(b))//8
	var i int = 0x18
	t.Log("i=", i)
	t.Log("&i=", &i)
	t.Log("unsafe.Pointer(&i)=", unsafe.Pointer(&i))
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	t.Log("bs=", bs)
	if bs[0] == 0 {
		fmt.Println("system edian is little endian")
	} else {
		fmt.Println("system edian is big endian")
	}
}

func TestBigEndian(t *testing.T) {

	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use big endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func TestLittleEndian(t *testing.T) {
	var i int = 0x12345678
	fmt.Println("十进制：", i)
	fmt.Println("16进制:", "0x12345678")
	fmt.Println("2进制:", DecBin(int64(i)))
	// 0000 0000 0000 0000   0000 0001 1111 1111
	var testInt int32 = 256
	fmt.Printf("%d use little endian: \n", testInt)
	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt))
	fmt.Println("int32 to bytes:", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes)
	fmt.Printf("bytes to int32: %d\n\n", convInt)
}

func DecBin(n int64) string {
	if n < 0 {
		log.Println("Decimal to binary error: the argument must be greater than zero.")
		return ""
	}
	if n == 0 {
		return "0"
	}
	s := ""
	for q := n; q > 0; q = q / 2 {
		m := q % 2
		s = fmt.Sprintf("%v%v", m, s)
	}
	return s
}
