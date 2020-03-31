package main

import "sort"
/*
	912. 排序数组
给定一个整数数组 nums，将该数组升序排列。



示例 1：

输入：[5,2,3,1]
输出：[1,2,3,5]
示例 2：

输入：[5,1,1,2,0,0]
输出：[0,0,1,1,2,5]
*/
//解法一：直接用sort函数，但是感觉不太对，太简单了
func sortArray(nums []int)[]int {
	sort.Ints(nums)
	return nums
}
//解法二：遍历
func sortArray1(nums []int)[]int{
	for i := 0 ; i <len(nums);i++{
		for j := i;j>0;j--{
			if nums[j]<nums[j-1]{
				nums[j],nums[j-1] = nums[j-1],nums[j]
			}
		}
	}
	return nums
}
func main(){

}