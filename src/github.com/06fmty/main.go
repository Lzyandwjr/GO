package main

import "fmt"

//fmt占位符
func main() {
	var n = 100
	fmt.Printf("%T\n", n) //查看字符类型
	fmt.Printf("%v\n", n) //布朗型
	fmt.Printf("%b\n", n) // 2进制
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n) // 8进制
	fmt.Printf("%x\n", n) //16进制

	var s = "hello"
	fmt.Printf("%s\n", s)
}
