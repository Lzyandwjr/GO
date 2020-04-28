package main

import "fmt"

// 空接口
// 接口中没有定义任何需要的实现方法时，该接口则为空接口
// 任意类型都实现了空接口，即是 空接口可以存储任意值
type xx interface {

}
// 空接口应用
// 1.空接口类型作为函数的参数
// 2.空接口可以作为map的value
//
func main(){
	var x interface{} // 定义一个空接口变量
	 x = 100
	 fmt.Println(x)
	 x = "hello"
	 fmt.Println(x)
	 x = false
	 fmt.Println(x)

	 //var m =  make(map[string]interface{},16)
	 //m["name"] =  "娜扎"
	 //m["age"]  = 18
	 //m["hobby"] = []string{"basketball","football","swimming"}
	 //fmt.Println(m)
	 ret,ok:= x.(int) //类型断言,猜的类型不对时，会返回一个布尔值
	 if !ok{
	 	fmt.Println("Type u guess is wrong")
	 }else {
	 	fmt.Println(ret)
	 }
	 fmt.Println(ret)
	 // 类型断言使用类外一种写法：switch
	switch x.(type) {
	case string:
		fmt.Println("字符串")
	case int:
		fmt.Println("整形")
	case bool:
		fmt.Println("布尔型")
	default:
		fmt.Println("猜不到")
	}
}