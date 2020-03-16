package main

import "fmt"

var name string
var age int
var isok bool

//批量声明
var (
	name1 string
	age2  int
	isokk bool
)

func main() {
	name = "lixiang"
	age = 16
	isok = true
	//var heihei string
	//go中非全局变量声明必须使用,不使用无法编译
	fmt.Printf("name:%s", name) //格式化
	fmt.Println(age)            //换行
	fmt.Print(age)              //在终端中输出要打印的内容

	s2 := "30"
	fmt.Println(s2)
}
