package main

import "fmt"

/*
有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，从而可以得到恰好 z升 的水？

如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。

你允许：

装满任意一个水壶
清空任意一个水壶
从一个水壶向另外一个水壶倒水，直到装满或者倒空
示例 1: (From the famous "Die Hard" example)

输入: x = 3, y = 5, z = 4
输出: True
*/
// 求x,y的最大公约数，并看其能否被z整除
func canMeasureWater(x int,y int ,z int) bool{
	if x+y < z {
		return false
	}
	if x == 0 || y == 0 {
		return z == 0 || x+y == z
	}
	return z%gcd(x,y) == 0
}
//求x,y的最大公约数
func gcd(a int,b int) int {
	for a%b!=0 {
		a,b =b,a%b
	}
	fmt.Println(b)
	return b
}
func main(){
	x := 3
	y := 5
	z := 4
	fmt.Println(canMeasureWater(x,y,z))
}
