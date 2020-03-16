package main

import "fmt"

/*
func main(){
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n",a,&a)
	fmt.Printf("b:%p type:%T\n",b,b)
	fmt.Println(&b)
	c := *b // 根据内存地址去取值
	fmt.Println(c)
}
*/

func modify1(x int){
	x = 100
}

func modify2(y *int){
	*y = 100
}

func main(){
	a :=1
	modify1(a)

	fmt.Println(a)

	modify2(&a)
	fmt.Println(a)

}