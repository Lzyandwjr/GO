package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	if num == 0 {
		return int(num)
	}
	var count int = 0
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			count++
		}
	}
	return count
	//return strings.Count(fmt.Sprintf("%d", num), "1")
}
func main() {
	var num uint32 = 00000010000
	fmt.Println(hammingWeight(num))
}
