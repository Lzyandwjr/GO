package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 找出数组中和为指定值的两个元素下表，比如数组{1,3,5,7,8}中和为8的下标为(0,3),(1,2)

func main() {
	array := [...]int{1, 3, 5, 7, 8}
	fmt.Printf("和为8的下标为： \n")
	for i := 0; i < len(array); i++ {
		for j := 0; j <= i; j++ {
			if array[i]+array[j] == 8 {
				fmt.Printf("(%v,%v) \n", j, i)
			}
		}
	}
	// for i,v1:=range array{
	// 	for j,v2 := range array{
	// 		if v1+v2 == 8{
	// 			fmt.Println(i,j)
	// 		}
	// 	}

	// }
	//创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。使用 for 循环访问所有元素并打印出来。提示:字符数据运算 'A'+1 -> 'B'
	eleArray := [26]byte{}

	for i := 0; i < 26; i++ {
		eleArray[i] = 'A' + byte(i)
		fmt.Printf("%c ", eleArray[i])
	}
	fmt.Println("\n")
	// 请求出一个数组的最大值，并得到对应的下标

	maxIndex := 0
	numberArray := [...]int{1, 3, -1, 10, -5, 30, 20, 5}
	maxValue := numberArray[0]
	for i := 0; i < len(numberArray); i++ {
		if maxValue < numberArray[i] {
			maxIndex = i
			maxValue = numberArray[i]
		}
	}
	fmt.Printf("maxValue:%v maxIndex:%v \n", maxValue, maxIndex)

	//)请求出一个数组的和和平均值
	sumArray := [...]int{1, 2, 3, 4, 78, 1, 2, 99, 53}
	sum := 0
	for i := range sumArray {
		sum = sum + sumArray[i]
	}
	//  fmt.Println(sum)
	sum = sum / len(sumArray)
	//  fmt.Println(len(sumArray))
	fmt.Println(sum)
	// 随机生成五个数，并将其反转打印

	var randArray [5]int
	//rand.Seed(time.Now().UnixNano()) // 设置种子，不然每次随机都会成0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(randArray); i++ {
		randArray[i] = rand.Intn(100)
	}
	fmt.Println("随机数组为:", randArray)

	temp := 0
	for i := 0; i < len(randArray)/2; i++ {
		temp = randArray[len(randArray)-i-1]
		randArray[len(randArray)-i-1] = randArray[i]
		randArray[i] = temp
	}
	fmt.Println("翻转数组为:", randArray)
}
