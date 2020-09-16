package main

import "fmt"

func printNumbers(n int) []int {
	var s []int
	m := int(pow(10, n))
	for i := 1; i < m; i++ {
		s = append(s, i)
	}
	return s
}

func pow(x float64, y int) float64 {
	res := 1.0
	for y != 0 {
		res *= x
		y--
	}
	return res
}

func main() {
	fmt.Println(printNumbers(3))
}
