package main

import "fmt"

func main(){
	finger := 3
	switch finger{
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	default:
		fmt.Println("无效输入")
	}
	// break
	// for i := 0;i < 5; i++{
	// 	fmt.Println(i)
	// 	if i == 3 {
	// 		break
	// 	}
	// }
	// continue 跳过本次循环，继续循环
	for i := 0;i<5;i++{
		fmt.Println(i)
		if i == 3{
			continue
		}
	}

	for j := 0;j<=10;j++{
		for k :=0;k<j;k++{
			if k == 3{
				goto breakTag
			}
			fmt.Printf("%v-%v\n",j,k)
		}
	}
	breakTag:
		fmt.Println("结束")
}