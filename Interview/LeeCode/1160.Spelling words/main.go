package main

import "fmt"

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