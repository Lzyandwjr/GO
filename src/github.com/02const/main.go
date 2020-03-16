package main

import "fmt"

//常量
const pi = 3.1415926

//批量声明常量
const (
	statusOK = 200
	nodFIND  = 404
)

//若在批量声明常量中某一行没有写值，则该常量=上面一个声明的常量
const (
	n1 = 100
	n2
	n3
)

//iota
const (
	a1 = iota
	a2
	a3
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2 //d1:1 d2:2
	d3, d4 = iota + 1, iota + 2 //d3:2 d4:3
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // << 左移 10位 2进制：1000000000 -> 十进制:1024
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(d1, d2, d3, d4)
	fmt.Println(KB, MB, GB, TB, PB)
}



