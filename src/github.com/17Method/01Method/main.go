package main

import "fmt"

// 方法定义实例

type Person struct {
	name string
	age int8
}
// newperson()是构造构造函数
func newPerson(name string,age int8) *Person{
	return &Person{
		name: name,
		age:  age,
	}
}

// Dream是为Person类型定义方法
func (p Person) Dream(){
	fmt.Printf("%s的梦想是学好Go语言\n",p.name)
}
// 指针接收者指的就是接收者的类型是指针
func (p *Person) SetAge(newAge int8){
	p.age = newAge
}
//使用值类型的接收者
func (p Person) SetAge2(newAge int8){
	p.age =  newAge
}
func main(){
	p1 := newPerson("刘哲源",int8(25))
	(*p1).Dream()
	p1.Dream()

	fmt.Println(p1.age)
	p1.SetAge(int8(18))
	fmt.Println(p1.age)
	(*p1).SetAge2(int8(6))
	fmt.Println(p1.age)
}
