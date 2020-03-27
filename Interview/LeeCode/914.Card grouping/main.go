package main

/*
给定一副牌，每张牌上都写着一个整数。

此时，你需要选定一个数字 X，使我们可以将整副牌按下述规则分成 1 组或更多组：

每组都有 X 张牌。
组内所有的牌上都写着相同的整数。
仅当你可选的 X >= 2 时返回 true。

 

示例 1：

输入：[1,2,3,4,4,3,2,1]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[3,3]，[4,4]
示例 2：

输入：[1,1,1,2,2,2,3,3]
输出：false
解释：没有满足要求的分组。
示例 3：

输入：[1]
输出：false
解释：没有满足要求的分组。
示例 4：

输入：[1,1]
输出：true
解释：可行的分组是 [1,1]
示例 5：

输入：[1,1,2,2,2,2]
输出：true
解释：可行的分组是 [1,1]，[2,2]，[2,2]

*/
// 思路： 首先计算数组中数值的个数，并求其最大公约数，如可以被len(deck)整除，则输出true
func hasGroupsSizeX(deck []int)bool{
	m := make(map[int]int,len(deck))
	for _,v := range deck{
		m[v]++
	}
	X := -1
	for _,v := range m{
		if X == -1{
			X =v
		}else {
			X = gcd(X,v)
		}
	}
	return X >=2
}

func gcd(x,y int)int{
	if y == 0{
		return x
	}
	return gcd(y,x%y)
}
func main(){

}
