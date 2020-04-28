package main

import "fmt"

// go语言中的接口是一种类型，抽象的类型


type Dog struct {

}
func (d Dog)say(){
	fmt.Println("汪汪汪")
}

type Cat struct {

}
func (c Cat)say(){
	fmt.Println("喵喵喵")
}
type Person struct {
	name string
}
func (p Person)say(){
	fmt.Println("啊啊啊")
}
// 接口不管是什么类型，他只管你要实现什么方法
//定义一个类型，一个抽象的累心，只要实现了say()这个方法类型都可以成为saye类型
type sayer interface {
	say()
}
// 打的函数
func da(arg sayer){
	arg.say()
}
func main(){
	//c1 := Cat{}
	//da(c1)
	//
	//d1 := Dog{}
	//da(d1)
	//
	//p1 := Person{name:"wujinran"}
	//
	//da(p1)

	var s sayer
	c2 := Cat{}
	s = c2
	p2 := Person{name:"sss"}
	s = p2
	fmt.Println(s)
}
