package main

import "fmt"

/*
	给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。

在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。

注意:
假设字符串的长度不会超过 1010。

示例 1:

输入:
"abccccdd"

输出:
7

解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
*/
// 回文串就是一个正读和反读都一样的字符串，比如“level”或者“noon”等等就是回文串。回文子串，顾名思义，即字符串中满足回文性质的子串。比如输入字符串 "google”，由于该字符串里最长的对称子字符串是 "goog”，因此输出4。
// 回文串就是对称
// 首先判断字符串长度是否超出值，未超出则继续，计算字符串中每个字符的个数，如果是双数则可以对称，如果是单数，则+1
func longestPalindrome(s string) int {
	if len(s) > 1010 {
		fmt.Println("字符串长度过长")
	}
	num := make(map[byte]int)
	count := 0
	single := 0
	for i := 0; i<len(s);i++{
		num[s[i]] += 1
	}
	for _,v := range num{
		count += v/2
		if v%2 == 1{
			single = 1
		}
	}
	return 2*count + single


}

func main(){
	var  s = "abccccdd"
	fmt.Println(longestPalindrome(s))
}