package main

import (
	"fmt"
	"sort"
)

/*
	输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。



示例 1：

输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
示例 2：

输入：arr = [0,1,2,1], k = 1
输出：[0]
*/


func getLeastNumbers(arr []int,k int) []int {
	if k > len(arr){
		fmt.Println("请重新输入")
	}else if k == 0{
		return []int{}
	}else if len(arr)>10000{
		fmt.Println("请重新输入")
	}else if k == len(arr){
		return arr
	}
	sort.Ints(arr[:])
	return arr[0:k]
}

func main(){
	arr := []int{8,6,10,9741,32,1}
	k := 5
	fmt.Println(getLeastNumbers(arr,k))
}