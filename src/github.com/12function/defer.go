package main

import "fmt"

// defer:延迟执行
func defer1(){
	fmt.Println("start...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end...")
}