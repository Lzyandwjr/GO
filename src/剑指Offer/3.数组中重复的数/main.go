package main

import "fmt"

/*
	找出数组中重复的数字。


在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

示例 1：

输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3

*/
//func FindNum(s []int) int {
//	sort.Ints(s)
//
//	for i := 0; i < len(s); i++ {
//		if s[i] == s[i+1] {
//			return s[i]
//		}
//	}
//	return -1
//}
//
//func main() {
//
//}
func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, value := range nums {
		fmt.Println(value)
		if _, ok := m[value]; ok {
			fmt.Println(ok)
			// 存在该数值
			return value
		} else {
			// 不存在
			m[value] = 1
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	num := findRepeatNumber(nums)
	fmt.Println("重复的一个数字为: ", num)
}
