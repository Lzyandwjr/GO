package main

import "fmt"

func min(x,y int) int{
	if x <= y{
		return x
	}
	return y
}

func updateMatrix(matrix [][]int) [][]int {

	m,n := len(matrix),len(matrix[0])
	fmt.Println(m)
	fmt.Println(n)
	dp := make([][]int,m)
	fmt.Println(dp)
	for i := 0;i<m;i++{
		dp[i] =  make([]int,n)
		fmt.Println(dp)
		for j:= 0;j<n;j++{
			if matrix[i][j] != 0{
				dp[i][j] = 10000
				if i != 0{
					fmt.Println(dp[i-1][j]+1)

					dp[i][j] = min(dp[i][j],dp[i-1][j]+1)
				}
				if j != 0 {
					dp[i][j] = min(dp[i][j],dp[i][j-1]+1)
				}
			}else {
				dp[i][j] = 0
			}
		}
	}
	for i := m-1;i >=0 ;i-- {
		for j := n - 1; j >= 0; j-- {
			if dp[i][j] > 1 {
				if i < m-1 {
					dp[i][j] = min(dp[i][j], dp[i+1][j]+1)
				}
				if j < n-1 {
					dp[i][j] = min(dp[i][j], dp[i][j+1]+1)
				}
			}
		}

	}
	return dp
}

func main(){
	dp := [][]int{{0,0,0},{0,1,0},{0,0,0},}
	updateMatrix(dp)
}
