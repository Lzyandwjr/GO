package main

import (
	"fmt"
	"reflect"
)

// reflect demo

func reflectType(x interface{}){
	// 此时不知道别人调用此函数时，传入进来什么类型的变量
	// 1.方式1：通过类型断言
	// 2.方法2：借助反射
	Obj := reflect.TypeOf(x)
	fmt.Println(Obj)
}

func reflectValue(x interface{}){
	v := reflect.ValueOf(x)
	fmt.Printf("%v,%T",v,v)
}

func reflectSetValue(x interface{}){
	v := reflect.ValueOf(x)
	// Elem()用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)

	}
}

type Cat struct {

}

type Dog struct {

}

func main(){
	//var a float32 = 1.234
	//reflectType(a)
	//var b int8 = 1
	//reflectType(b)
	//var c Cat
	//var d Dog
	//reflectType(c)
	//reflectType(d)
	//// slice
	// var e []int
	//var f []string
	//reflectType(e)
	//reflectType(f)

	// value of
	var aa int32 = 100
	reflectValue(aa)

	// set value
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)
 }
