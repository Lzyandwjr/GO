package main

import "fmt"

func main(){
	var a []string
	var b []string


	var c  = []bool{false,true}
	fmt.Println(a,b,c)
	// 基于数组得到切片
	d := [5]int{1,7,6,3,7}
	f := d[1:4]
	fmt.Println(d,f)
	fmt.Printf("%T \n",f)

	//切片再次切片

	e := f[0:len(f)]
	fmt.Printf("%T\n",e)

	// make函数构造切片
	g := make([]int,5,10)
	fmt.Println(g)
	fmt.Printf("%T\n",g)

	//通过len函数获取切片 长度
	fmt.Println(len(g))
	// 通过cap()函数获取切片容量
	fmt.Println(cap(g))

	//var i []int //声明int类型切片
	//var j = []int{} //声明并且初始化
	//if i == nil {
	//	fmt.Println("i是一个nil")
	//}
	//if j == nil {
	//	fmt.Println("j是一个nil")
	//}
	//fmt.Println(i,len(i),cap(i))
	//fmt.Println(j,len(j),cap(j))

	// 切片的数只拷贝
	 i := make([]int,3)
	 j := i
	 j[0] = 100
	 fmt.Println(i)
	 fmt.Println(j)

	 // 切片的遍历

	 s := []int{1,2,3,4,5}
	 for i := 0 ;i <len(s);i++{
	 	fmt.Println(i,s[i])
	 }
	 fmt.Println()
	 for index,value := range s{
	 	fmt.Println(index,value)
	 }

	 // 切片扩种
	  var z []int
	  for i :=0 ;i<=10;i++{
	  	z = append(z,i)
		  fmt.Println(z,len(z),cap(z))
	  }
	  z = append(z,11,53,55,7,8)
	  fmt.Println(z)
	  x := s[0:3]
	  z = append(z,x...)
	fmt.Println(z)
	// 切片复制
	  h :=  make([]int,20)
	  copy(h,z)
	  h[0]= 666
	  fmt.Println(h)

	  // 切片删除元素
	  h = append(h[0:2],h[3:]...)
	  fmt.Println(h)
}