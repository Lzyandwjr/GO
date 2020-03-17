package main

import "fmt"

//结构体初始化
// 1.键值对初始化
//2.值的列表进行初始化
type per1s struct {
	name,city string
	age int
}

func main(){
	//键值对初始化

	p1 := per1s{
		name: "lll",
		city: "111",
		age:  2,
	}
	fmt.Printf("%#v\n",p1)

	p5 := &per1s{
		name: "asda",
		city: "1dasd",
		age:  10,
	}
	fmt.Printf("%#v\n",p5)
	//2.值得列表进行初始化 , 必须按照结构体的顺序来,且必须初始虎啊结构体的所有字段，不能与键值初始化方式混用
	p6 := per1s{
		 "kaa",
		 "km",
	  12,
	}
	fmt.Printf("%#v\n",p6)
}