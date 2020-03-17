package main
import "fmt"

//构造函数：构造一个结构体实例的函数
// 结构体是值类型，
type person struct {
	name,city string
	age 	  int8
}
//下面此法性能较差
//func newPerson(name,city string,age int8) person{
//	return person{
//		name: name,
//		city: city,
//		age:  age,
//	}
//}
func newPerson(name,city string,age int8) *person{
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main(){
	p1 := newPerson("lll","shshs",int8(18))
	fmt.Println(p1)
}