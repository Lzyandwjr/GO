package main

import "fmt"

/*
	给你一份『词汇表』（字符串数组） words 和一张『字母表』（字符串） chars。

假如你可以用 chars 中的『字母』（字符）拼写出 words 中的某个『单词』（字符串），那么我们就认为你掌握了这个单词。

注意：每次拼写时，chars 中的每个字母都只能用一次。

返回词汇表 words 中你掌握的所有单词的 长度之和。



示例 1：

输入：words = ["cat","bt","hat","tree"], chars = "atach"
输出：6
解释：
可以形成字符串 "cat" 和 "hat"，所以答案是 3 + 3 = 6。
*/
func countCharacters(words []string,chars string) int {
	count := 0
	for _,value := range words {
		freq_count := make(map[byte]int)
		for i := 0;i < len(chars);i++{
			freq_count[chars[i]] = freq_count[chars[i]]+1
		}
		flag := true
		for i := 0; i<len(value);i++{
			if freq_count[value[i]] > 0{
				freq_count[value[i]]--
			}else if freq_count[value[i]] <= 0 {
				flag = false
				break

			}
		}
		if flag {
			count += len(value)
		}
	}
	return count
}

func main(){
	var words = []string{"words","shy","happy","apple","app"}
	var chars = "apleshy"
	fmt.Printf("%T",words)
	count := countCharacters(words,chars)

	fmt.Printf("字母长度为：%d",count)
}