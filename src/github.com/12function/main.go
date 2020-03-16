package main

import "fmt"




func sayHello(){
	fmt.Println("hello it's me")
}

func  Name(name string){
	fmt.Println("hhhh",name)
}

func jisuan(i int,j int)int{
	return i+j
}

func main(){
	sayHello()

	//i := "吴尽然"
	//Name(i)
    //k :=1
    //s :=2
    //fmt.Println(jisuan(k,s))
    //r1 := intSum2(0)
    //r2 := intSum2(10)
    //r3 := intSum2(10,20)
    //r4 := intSum2(10,20,30)
    //fmt.Println(r1)
    //fmt.Println(r2)
    //fmt.Println(r3)
    //fmt.Println(r4)

    x,y := calc(100,200)
    fmt.Println(x,y)

    defer1()

}

// 函数接受可变参数,在参数后面+... 标识可变参数。可变参数在函数体中是切片
func intSum(a ...int) int {
	ret := 0
	for _,arg := range a{
		ret = ret + arg
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum2(a int,b ...int)int{
	ret := a
	for _,arg := range b{
		ret = ret +arg
	}
	return ret
}

// 多个返回值的函数
func calc(a,b int)(sum,sub int){
	sum = a+b
	sub = a-b
	return sum,sub
}