package main

import (
	"fmt"
	"sort"
)

func main(){
	// 使用内置sort包对数组a{3,7,8,9,1}进行排序
	var a =[...]int{3,7,8,9,1}
	// a[:]得到的是一个切片
	sort.Ints(a[:])
	fmt.Println(a)
}
