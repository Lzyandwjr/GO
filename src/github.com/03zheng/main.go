package main

import "fmt"

func main() {
	//十进制
	var i1 = 101
	fmt.Printf("%d \n", i1)
	fmt.Printf("%b \n", i1) //转2进制
	fmt.Printf("%o \n", i1) //十进制 转换成8进制
	fmt.Printf("%x \n", i1) //十进制转换成16进制

	// 8进制
	i2 := 077
	fmt.Printf("%d", i2)

	//16进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	//查看变量类型
	fmt.Printf("%T \n", i3)

	//声明int8类型的变量
	i4 := int8(9)
	fmt.Printf("%T \n", i4)
}


