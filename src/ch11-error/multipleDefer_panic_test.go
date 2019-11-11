package ch11_error

import (
	"errors"
	"fmt"
	"testing"
)

//多个defer，defer抛出恐慌，下个defer捕获恐慌
func TestMultipleDefer(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("panic:", e)
		}
	}()

	defer func() {
		panic(errors.New("error from defer"))
	}()
}
