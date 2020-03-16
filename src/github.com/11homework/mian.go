package main

import (
	"fmt"
	"strings"
)

func main(){
	// 统计一个字符串中每个单词出现的次数
	// "how do you do“中每个单词出现的次数

	//0.定义一个map[string]int
	var s = "how do you do"
	var wordCount =  make(map[string]int,10)
	//1.字符串中都有哪些单词
	words := strings.Split(s," ")
	//2.遍历单词做统计
	for _,word := range words{
		v,ok := wordCount[word]
		if ok {
			wordCount[word] = v + 1
		}else {
			wordCount[word] = 1
		}
	}
	for k,v := range wordCount{
		fmt.Println(k,v)
	}
}
