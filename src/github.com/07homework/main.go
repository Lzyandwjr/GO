package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "helloworld沙河小王子"
	numcount := 0

	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			numcount++
		}
	}
	fmt.Printf("汉字个数为：%v\n", numcount)
}
