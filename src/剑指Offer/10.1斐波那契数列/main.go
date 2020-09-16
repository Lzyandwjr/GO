package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	sum := []int{0, 1}
	for i := 2; i <= n; i++ {
		sum = append(sum, sum[i-1]+sum[i-2])
	}
	return sum[n]

}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = (dp[i-2] + dp[i-1]) % (1e9 + 7)
	}
	return dp[n]
}

func main() {
	fmt.Println(fib(10))
	fmt.Println(fib2(10))
}
