package  main

import (
	"fmt"
	"strings"
)

// 匿名函数和闭包
//定义一个函数它的返回值是一个函数
//把函数作为返回值
func a(name string) func() {
	return func(){
		fmt.Println("hello 666",name)
	}
}

func  main(){
	// 闭包 = 函数+外层变量的引用
	r := makeSuffixFunc(".txt")
	ret := r("lzy")
	fmt.Println(ret)

	x,y := calc(100)
	ret1 := x(200) // base = 100+200=300
	ret2 := y(300)	// base= 300-300
	fmt.Println(ret1)
	fmt.Println(ret2)
}


func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string{
		if !strings.HasPrefix(name,suffix) {
			return name + suffix
		}
		return name
	}
} // 使用闭包判断文件拓展名


func calc(base int) (func(int) int,func(int) int){
	add :=func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add,sub
}