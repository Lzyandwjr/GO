package main

import (
	"fmt"
	"math"
	"strings"
)

var pyramid [][]int = [][]int{
	{1},
	{2, 3},
	{4, 5, 6},
	{7, 8, 9, 10},
}

func FindBestWay(pyramid [][]int) int {
	row := len(pyramid)
	for i := row - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			maxNum := int(math.Max(float64(pyramid[i][j]), float64(pyramid[i][j+1])))
			pyramid[i-1][j] += maxNum
		}
	}
	return pyramid[0][0]
}

func main() {
	//var pyramid [][]int
	//var n int
	//fmt.Scanf("%d", &n)
	//for i := 1; i <= n; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Scanln(&pyramid[i][j])
	//	}
	//}
	fmt.Println(FindBestWay(pyramid))
	fmt.Println(strings.ToUpper("Application of improved Particle swarm optimization algorithm in container Scheduling strategy"))
}
