package main

import "fmt"

// 数组相关
func main(){
	var a [3]int
	var b [4]int
	fmt.Println(a)
	fmt.Println(b)

	//数组初始化
	// 1.定义时使用初始值列表的方式初始化
	var cityArray = [4]string{"Peking","Shanghai","Shenzhen","Guangzhou"}
	fmt.Println(cityArray)
	//2. 编译器推导长度
	// var boolArray = [...]bool{true,false,false,true,false}
	// fmt.Println(len(boolArray))
	// 3. 使用索引值方式初始化
	var langArray = [...]string{1:"go",3:"python",7:"java"}
	fmt.Println(langArray)
	fmt.Println(len(langArray))
	fmt.Printf("%T\n",langArray)


	// 数组的遍历
	for i :=0;i<len(cityArray);i++{
		fmt.Println(cityArray[i])
	}
	for index, value := range cityArray{
		fmt.Println(index,value)
	}
	for _, value1 := range cityArray{
		fmt.Println(value1)
	}
	for index := range cityArray{
		fmt.Println(index)
	}

	// 二维数组
//多维数组只有只有最外层可以使用简写方式
	 city := [...][2]string{
		 {"peking","xi'an"},
		 {"Shanghai","Chongqing"},
		 {"Guangzhou","Shenzhen"},
		 {"Chengdu","Guiyang"},
	 }
	 fmt.Println(city[2][1])

	 for _,v1 := range city{
		 for _,v2 := range v1{
			 fmt.Println(v2)
		 }
	 }
	 //多维数组只有只有最外层可以使用简写方式

	 // 数组是值类型
	 x := [3]int{1,2,3}
	 fmt.Println(x)
	 f1(x)
	 fmt.Println(x)
}

func f1(a [3]int){
	a[0]=100
}