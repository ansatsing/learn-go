package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math/big"
	"testing"
)

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
func main() {
	data1 := []byte("I like donuts")
	data2 := []byte("I like donutsca07ca")
	targetBits := 4
	target := big.NewInt(1)
	fmt.Println(target)
	target.Lsh(target, uint(256-targetBits))
	fmt.Printf("%x\n", sha256.Sum256(data1))
	fmt.Printf("%64x\n", target)
	fmt.Printf("%x\n", sha256.Sum256(data2))

	st := "10000000000000000000000000000000000000000000000000000000000"
	fmt.Println(256 / len(st))
	var hashInt big.Int
	counter := 0
	for {
		data1 = bytes.Join([][]byte{data1, IntToHex(int64(counter))}, []byte{})
		has := sha256.Sum256(data1)
		fmt.Printf("runing:%x\n", has)
		hashInt.SetBytes(has[:])
		if hashInt.Cmp(target) == -1 {
			fmt.Println()
			fmt.Printf("result:%x\n,counter:%x\n", has, counter)
			break
		} else {
			counter++
		}
	}
}

func TestSha(t *testing.T) {
	str := "ansatsing"
	hsh := sha256.Sum256([]byte(str))
	fmt.Printf("%x\n,%s\n,%d\n", hsh, hsh, hsh)
}
