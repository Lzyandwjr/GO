package main

import "fmt"

func myPow(x float64, n int) float64 {
	/*
		超出时间限制
	*/
	//var tmp float64 = 1
	//
	//if n == 0 {
	//	return float64(1)
	//}
	//if n == 1 {
	//	return x
	//}
	//for n != 0 {
	//	if n > 0 {
	//		for i := 0; i < n; i++ {
	//			tmp = tmp * x
	//		}
	//		return tmp
	//	} else {
	//		m := -n
	//		for i := 0; i < m; i++ {
	//			tmp = tmp * x
	//		}
	//		return 1 / tmp
	//	}
	//
	//}
	//return float64(tmp)

	var tmp float64 = 1.0
	if n == 0 {
		return 1.0
	} else if n < 0 {
		x = 1.0 / x
		n = -n
	}
	if n == 1 {
		return x
	}
	for n > 1 {
		//判断n是否为奇数（判断n的二进制最后一位是不是1）
		if n&1 == 1 {
			tmp *= x
			n--
		} else {
			//如果为偶数，就能够直接移位
			x *= x
			n = n >> 1
		}
	}
	tmp *= x
	return tmp

}
func main() {
	fmt.Println(myPow(2.0000, 10))
}
