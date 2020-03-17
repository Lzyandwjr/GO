package main

import "fmt"

type  person struct {
	name string
	age int
	city string
}
// 指针类型结构体
type people struct {
	name1,city1 string
	age int8
}
func main(){
	var p1 person
	p1.name = "刘哲源"
	p1.age = 25
	p1.city = "丰县"
	fmt.Println(p1)


	// 匿名结构体
	var user struct{
		name string
		married bool
	}
	user.name = "吴尽然"
	user.married = true
	fmt.Printf("%#v\n",user)

	var p2 = new(people)
	fmt.Printf("%T\n",p2)
	(*p2).name1 = "lzy"
	(*p2).city1 = "GY"
	(*p2).age = 28
	// 指针类型的结构体，可以不用(*),直接.就可以
	var p3 = new(people)
	p3.age=30
	p3.name1="www"
	p3.city1 = "beijing"
	fmt.Println(p2)
	fmt.Println(p3)


	// 取结构体的地址进行实例化
	p4 := &people{}
	fmt.Printf("%T\n %#v\n",p4,p4)
	p4.city1 = "Nanjing"
	p4.name1 = "LLLL"
	p4.age = 10
	fmt.Println(p4)

}
