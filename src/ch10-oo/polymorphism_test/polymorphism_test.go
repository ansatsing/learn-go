package polymorphism_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	helloWorld() Code
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) helloWorld() Code {
	return "System.out.println(\"hello world!\")"
}

type GoProgrammer struct {
}

func (g *GoProgrammer) helloWorld() Code {
	return "fmt.println(\"hello world！\")"
}

type Company struct {
	p Programmer //不能用指针:p *Programmer
}

func (c *Company) writeCode() {
	fmt.Println(c.p.helloWorld())
}

func writeHelloWorldCode(p Programmer) { //不能用指针: p *Programmer
	fmt.Println(p.helloWorld())
}
func TestPolymorphism(t *testing.T) {
	jp := new(JavaProgrammer)
	gp := new(GoProgrammer)
	writeHelloWorldCode(jp)
	writeHelloWorldCode(gp)

	//编译报错：cannot use jp1 (type JavaProgrammer) as type Programmer in argument to writeHelloWorldCode:
	//	JavaProgrammer does not implement Programmer (helloWorld method has pointer receiver)
	//jp1:=JavaProgrammer{}
	//writeHelloWorldCode(jp1)

	cpny := Company{jp}
	cpny.writeCode()

	cpny.p = gp
	cpny.writeCode()
}
