package main

func superEggDrop(K,N int)int{
	dp := make([][]int, K + 1)
	for i := 0; i <= K; i++ {
		dp[i] = make([]int, N + 1)
	}

	for j := 1; j <= N; j++ {
		for i := 1; i <= K; i++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j-1] + 1
			if dp[i][j] >= N {
				return j
			}
		}
	}
	return N
}

func main(){

}
