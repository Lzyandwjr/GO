package main

import "fmt"


// 求数组{1,3,5,7,8}所有元素之和
func main(){
	s := 0
	a := [...]int{1,3,5,7,8}
	for i := range a{
		s = a[i] + s 
	}
	fmt.Println(s)
}