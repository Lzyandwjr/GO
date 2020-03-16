package main

import (
	"fmt"
)
// if条件判断
func main(){
	age := 19
	if age > 18{
		fmt.Printf("great \n")
	}else{
		fmt.Printf("fuck \n")
	}

	if age > 35{
		fmt.Printf("人到中年 \n")
	}else if age > 18{
		fmt.Printf("精神小伙 \n")
	}else{
		fmt.Printf("中二少年 \n")
	}
	//for循环
	var i = 0
	for  ;i<10;i++{
		fmt.Println(i)
	}

	// for 1
	for i := 0; i<10;i++{
		fmt.Println(i)
	}
	// for 2
	for ;i<10; {
		fmt.Println(i)
		i++
	}

	// for range 循环
	s := "hello world"
	for a,b :=range s{
		fmt.Printf("%d %c \n", a, b)
	}
}