package ch15_objectPool

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	connPool := NewConnectionPool(2, 10)
	con, err := connPool.getConnection(time.Millisecond * 300)
	if err == nil {
		fmt.Println(con)
		fmt.Println(connPool.len)
	}
	con, err = connPool.getConnection(time.Millisecond * 300)
	if err == nil {
		fmt.Println(con)
		fmt.Println(connPool.len)
	} else {
		fmt.Println(connPool.len)
		fmt.Println(con)
		fmt.Println(err)
	}

}

func TestStrconv(t *testing.T) {
	var i int64
	i = time.Now().Unix()
	fmt.Println("i=", i)
	fmt.Println("base=", 10, " ", strconv.FormatInt(i, 10))
	fmt.Println("base=", 2, " ", strconv.FormatInt(i, 2))
	fmt.Println("base=", 3, " ", strconv.FormatInt(i, 3))
}
