package main

import (
	"fmt"
	"reflect"
)

// struct reflect

type student struct {
	Name string
	Grade int
}

func main(){
	stu1 := student{
		Name:  "lzy",
		Grade: 90,
	}
	//通过反射去获取结构体中所有字段的信息
	t:=reflect.TypeOf(stu1)
	fmt.Printf("name:%v,kind:%v\n",t.Name(),t.Kind())
	// 遍历结构体变量的所有字段
	for i:=0;i<t.NumField();i++{
	// 根据结构体字段的索引去取字段
		fileObj:=t.Field(i)
		fmt.Printf("name:%v,type:%v,tag:%v\n",fileObj.Name,fileObj.Type,fileObj.Tag)
	}
}