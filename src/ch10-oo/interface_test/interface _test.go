package interface_test

import "testing"

type Pgrogrammer interface {
	helloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) helloWorld() string {
	return "fmt.println(\"hello world\")"
}

func TestInterface(t *testing.T) {
	var programmer Pgrogrammer
	programmer = new(GoProgrammer)
	t.Log(programmer.helloWorld())
}
