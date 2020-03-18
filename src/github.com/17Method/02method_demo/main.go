package main

import "fmt"

// 为任意类型添加方法

type MyInt int

func (i MyInt) sayHi(){
	fmt.Println("hi")
}

func main() {
	var m1 MyInt
	m1 = 100
	m1.sayHi()
}