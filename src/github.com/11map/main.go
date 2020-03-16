package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// map()映射
func main(){
	var a map[string]int
	fmt.Println(a == nil)

	// map初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)
	// map中添加键值对
	a["吴尽然"] = 9999
	a["刘哲源"] = 1111
	fmt.Println(a)
	fmt.Printf("%T \n",a)

	b := map[int]bool{
		1:true,
		2:false,
		3:true,
		4:true,
	}
	fmt.Println(b)

	var scoreMap = make(map[string]int,8)
	scoreMap["刘哲源"] = 100
	scoreMap["吴尽然"] = 200
	//判断土豆在不在scoreMap中
	value,ok:= scoreMap["土豆"]
	fmt.Println(value,ok)
	if ok {
		fmt.Println("土豆在map中")
	}else {
		fmt.Println("查无此人")
	}
	// map是无序的
	for k,v := range scoreMap{
		fmt.Println(k,v)
	}
	// 只遍历map中的key
	for k:=range scoreMap{
		fmt.Println(k)
	}
	//	只遍历map中的value
	for _,v:=range scoreMap{
		fmt.Println(v)
	}
	// 删除map中的键值对
	delete(scoreMap,"吴尽然")
	fmt.Println(scoreMap)

	// 按照某个固定顺序遍历map
	var cityMap = make(map[string]int,100)
	// 添加50个键值对
	for i:=0;i<50;i++{
		key :=fmt.Sprintf("stu%02d",i)
		value := rand.Intn(100)
		cityMap[key] = value
	}
	for k,v := range cityMap{
		fmt.Println(k,v)
	}
	//按照key从小到大遍历map
	// 取出所有的key存放到切片中，对key做排序，按照排序后的key对citymap排序

	keys := make([]string,0,100)
	for k:=range cityMap{
		keys = append(keys,k)
	}
	//fmt.Println(keys)
	sort.Strings(keys)
	for _,key := range keys{
		fmt.Println(keys,cityMap[key])


	}
	// 元素类型为map的切
	var mapSlice = make([]map[string]int,8,8)// 完成了切片的初始化
	// [nil nil nil nil nil nil nil nil]
	// 还需要完成内部map元素的初始化
	mapSlice[0] = make(map[string]int,8)
	mapSlice[0]["吴尽然"] = 1000
	fmt.Println(mapSlice)

	// 值为切片的map
	var sliceMap = make(map[string][]int,8)//只完成了map的初始化
	v,ok := sliceMap["中国"]
	if ok{
		fmt.Println(sliceMap["中国"],v)
	}else {
		sliceMap["中国"] = make([]int,8,8)//完成了对切片的初始化
		sliceMap["中国"][0] = 100
		sliceMap["中国"][3] = 300
		sliceMap["中国"][2] = 200
		sliceMap["中国"][4] = 400
	}
	for k,v := range sliceMap{
		fmt.Println(k,v)
	}
}
