package ch20_io

import (
	"encoding/binary"
	"fmt"
	"os"
	"testing"
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
