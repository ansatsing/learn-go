package inherit_test

import (
	"fmt"
	"testing"
)

//复合
//Go 不不⽀支持继承，但可以通过复合的⽅方式来复⽤用

type Pet struct{}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println("   ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	d.p.Speak()
	fmt.Println("wang!")
}

func (d *Dog) SpeakTo(host string) {
	d.Speak()
	fmt.Println("   ", host)
}

func TestInheritByComposite(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("nba")
}

type Pig struct {
	Pet
	name string
}

//自己有就运行自己的，否则就运行Pet的
//func (d *Pig) Speak()  {
//	fmt.Println("zhuzhu!")
//}
//
//func (d *Pig) SpeakTo(host string)  {
//	d.Speak()
//	fmt.Println("   ",host)
//}

func TestInheritByOther(t *testing.T) {
	pig := new(Pig)
	pig.SpeakTo("xxx")

	pig.Speak()
}
