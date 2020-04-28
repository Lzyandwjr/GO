package main

import "fmt"

type Action interface {
	move()
	say()
}

// 使用值接收者实现接口和使用指针接收者实现接口的区别
type Mover interface {
	move()
}

type sayer interface {
	say()
}
type person struct {
	name string
	age int
}
// 使用值接收者实现接口： 类型的值和类型的指针都能够保运到接口变量中
func (p person)move(){
	fmt.Printf("%s is running\n",p.name)
}
// 使用指针接受则接收者实现接口：只有类型指针能够保存到接口变量中
func (p *person)say(){
	fmt.Printf("%s is saying\n",p.name)
}
func main(){
	var m Mover
	p1 := person{
		name:"wujinran",
		age:18,
	}

	m = p1
	m.move()
	fmt.Println(m)

	p2 := &person{
		name: "liuzheyuan",
		age:  20,
	}
	var n sayer
	n = p2
	n.say()
	fmt.Println(n)

	var a Action
	a = p2
	a.move()
	a.say()
}