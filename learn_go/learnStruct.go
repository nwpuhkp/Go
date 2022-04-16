package main

import "fmt"

type Info struct {
	Name       string
	Age        int
	University string
	mima       int
}

//方法
type MyInt struct {
	Number int
}

func (m MyInt) SayHello() {
	fmt.Println("Hello world")
}
func (m *MyInt) SetNumber(other int) { //若要修改结构体内的值需要传入指针
	m.Number = other
}
func (m MyInt) SayNumber() {
	fmt.Println(m.Number)
}

//组合(继承)
type ViewName struct {
	Name string
	ViewOther
}

func (v ViewName) SayName() {
	fmt.Println(v.Name)
}

type ViewOther struct {
	Value string
}

func (v ViewOther) SayValue() {
	fmt.Println(v.Value)
}
func (v *ViewOther) ChangeValue(change string) {
	v.Value = change
}

func main() {
	var info Info
	info = Info{
		Name:       "HuangKePu",
		Age:        21,
		University: "NWPU_cyber",
		mima:       123456,
	}
	var infoTwo = new(Info)
	//var infoTwo Info
	infoTwo.Name = "ZhuJiaWen"
	infoTwo.Age = 20
	infoTwo.University = "NWPU_software"
	fmt.Println(info, "\n", infoTwo, *infoTwo)
	fmt.Println(info.mima)

	var my MyInt
	my.Number = 1
	my.SayHello()
	my.SayNumber()
	my.SetNumber(100)
	my.SayNumber()

	var viewName ViewName
	viewName.Name = "Huangkepu"
	viewName.Value = "B"
	viewName.SayName()
	viewName.SayValue()
	viewName.ChangeValue("A")
	viewName.SayValue()
}
