package ch20_io

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestReaderFromFile(t *testing.T) {
	const filePath = "readfromfile_test.go"

	f, err := os.Open(filePath)

	defer func() { f.Close() }()

	if err != nil {
		t.Error("open file with problems")
	}
	buf := make([]byte, 8)
	n, err := f.Read(buf)
	if err != nil {
		t.Error("file read with problems")
	}
	if n > 0 {
		fmt.Println(buf)
		fmt.Println(uint64(buf[0]))
		v := int(binary.BigEndian.Uint64(buf))
		fmt.Printf("%d\n", v)
	}
}

func TestReadInt64FileByChunck(t *testing.T) {
	a := []uint64{34, 22, 78, 2, 1, 12, 87, 23}
	f, _ := os.Create("test.txt")
	buf := make([]byte, 8)
	for _, v := range a {
		binary.BigEndian.PutUint64(buf, uint64(v))
		f.Write(buf)
	}
	f.Close()
	time.Sleep(time.Millisecond * 100)
	rf, _ := os.Open("test.txt")
	inf, _ := rf.Stat()
	fmt.Printf("size:%d\n", inf.Size())
	rf.Seek(16, 0)
	for {
		n, _ := rf.Read(buf)
		if n <= 0 {
			break
		}
		fmt.Println(binary.BigEndian.Uint64(buf))
	}
}
