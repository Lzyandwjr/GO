package main

import "fmt"

// 1.自定义类型
// MyInt 基于int类型的自定义类型,类型为：main.MyInt，其中main为包名，value:0
type MyInt int
//2.类型别名：类似本名和英文名和小名的意思，其实为同一类型,下为NewInt其实就为int
type NewInt = int

func main(){
	var i MyInt
	fmt.Printf("%T value:%v \n",i ,i)
	var x NewInt
	fmt.Printf("%T,%v\n",x,x)
}