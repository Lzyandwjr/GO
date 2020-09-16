package main

import "fmt"

func shuiXianHua() int {
	for n := 100; n < 1000; n++ {
		a := n / 100
		b := (n - a*100) / 10
		c := n % 10
		if pow(a, 3)+pow(b, 3)+pow(c, 3) == n {
			return n
		}
	}
	return 0

}

func pow(x, y int) int {
	res := 1
	for y != 0 {
		res *= x
		y--
	}
	return res
}

func main() {
	for n := 100; n < 1000; n++ {
		a := n / 100
		b := (n - a*100) / 10
		c := n % 10
		if pow(a, 3)+pow(b, 3)+pow(c, 3) == n {
			fmt.Println(n)
		}
	}
}
