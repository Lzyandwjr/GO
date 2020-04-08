package main

func movingCount(m,n,k int)int{
	visit := make([][]bool,m+1)
	for i := range visit{
		visit[i] = make([]bool,n+1)
	}
	return DFS(0,0,m,n,k,visit)
}

func getSum(m,n int) int {
 	sum := 0
 	for m!=0{
 		sum+=m%10
 		m=m/10
	}
	for n!=0{
		sum+=n%10
		n=n/10
	}
	return sum
}

func DFS(i,j,m,n,k int,visit [][]bool)int{
	// 先检查i,j是否在合法范围内
	// 或者i,j这个位置已经访问过
	// 或者虽然没有访问过，但这个位置的数位之和大于k
	if i < 0 ||i >= m || j<0||j >=n||visit[i][j]||getSum(i,j) >k{
		return 0// 说明不能移动到这个位置
	}
	// 如果不是上面的情况,说明可以移动到i,j这个位置
	visit[i][j] = true// 于是把位置i,j标识为已访问过
	sum := 1
	sum += DFS(i-1,j,m,n,k,visit)//向上
	sum += DFS(i+1,j,m,n,k,visit)//向下
	sum += DFS(i,j-1,m,n,k,visit)//向左
	sum += DFS(i,j+1,m,n,k,visit)//向右
	return sum
}

func main(){}
