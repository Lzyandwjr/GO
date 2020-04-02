package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段的可见性与JSON序列化
// 如果一个Go语言包中定义的标识符是首字母大写的，那么对外可见的 如果小写则是对外不可见
type student struct{
	ID int //若首字母为小写 则对外不可见
	Name string
}

func newStudent(id int,name string)student{
	return student{
		ID:   id,
		Name: name,
	}
}
type class struct {
	Title string
	Students []student
}

func main(){
	//创建一个班级c1
	c1 := class{
		Title:  "冰糖炖雪梨",
		Students: make([]student,0,20),
	}
	// 往班级c1中添加学生
	for i := 0;i < 10;i++{
		tempstu := newStudent(i,fmt.Sprintf("stu%02d",i))
		c1.Students = append(c1.Students,tempstu)
	}
	fmt.Println(c1)
	// json序列化：go语言中的数据 -> json格式的字符串
	data,err := json.Marshal(c1)
	if err != nil{
		fmt.Println("json marshal failed ,err",err)
		return
	}
	fmt.Printf("%T\n",data)
	fmt.Printf("%s\n",data)
	// JSON反序列化：JSON格式的字符串 -> Go语言中的数据
	//jsonStr := ''
	//json.Unmarshal(jsonStr)

}
