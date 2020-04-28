package main

import (
	"fmt"
	"strings"
)

/*给定一个字符串，逐个翻转字符串中的每个单词。

 

示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

func reverWords(s string)string{
	s = strings.TrimSpace(s)
	arr := strings.Split(s," ")
	words := make([]string,0,len(arr))
	for i := len(arr)-1;i >=0;i--{
		fmt.Println(arr[i])
		if len(arr[i])!=0 {
			words = append(words, arr[i])
			fmt.Println(words)
		}
	}
	return strings.Join(words," ")
}

func main(){
	s :="  the sky is blue"
	fmt.Println(reverWords(s))
}