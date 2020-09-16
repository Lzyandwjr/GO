package main

import "fmt"

//题目描述
//在一个二维数组中（每个一维数组的长度相同），每一行都按照从左到右递增的顺序排序，
//每一列都按照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func Find(target int, array [][]int) bool {
	a := 0
	b := len(array[0]) - 1
	for {
		if a < len(array) || b >= 0 {
			v := array[a][b]
			if target == int(v) {
				return true
			} else if target < int(v) {
				b--
			} else {
				a++
			}
		}
		return false
	}
}

func main() {
	target := 5
	var array [][]int
	k := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			k++
			array[i][j] = k
		}
	}
	fmt.Println(array)
	fmt.Println(Find(target, array))
}
