package main

import(
	"fmt"
	"strings"
)


func main(){
	s := "hello world"

	//a1 := 'h'
	//a2 := 'e'
	//a3 := 'l'
	//a4 := 'o'
	fmt.Printf(s)
	//fmt.Printf("%s,%s,%s,%s,%s",a1,a2,a3,a3,a4)
	// \本来是具有特殊含义的应该告诉程序\只是单纯的\
	path := "\"F:\\Go\\src\\github.com\\07char\""
	fmt.Printf(path)

	s1 := "I'm ok"
	fmt.Printf(s1)

	// 多行字符串
	s2 := `
		世情薄
		人情恶
		雨送黄昏花易落
		`
	fmt.Printf(s2)

	s3 := `F:\Go\src\github.com\07char`
	fmt.Printf(s3)
	fmt.Print(len(s3))

	//字符串拼接
	name := "啊啊啊"
	world := "噢噢噢"
	ss := name + world
	fmt.Printf(name + world)
	fmt.Printf(ss)
	ss1 := fmt.Sprintf("%s%s",name,world)
	fmt.Printf(ss1)
	fmt.Printf("%s%s",name,world)

	//字符串分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(ss, "啊啊啊"))

	// 前缀
	fmt.Println(strings.HasPrefix(ss, "啊啊啊"))
	//后缀
	fmt.Println(strings.HasSuffix(ss, "噢"))
	//字符串出现的位置
	s4 := "abcde"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4,"d"))

	//join操作
	fmt.Println(strings.Join(ret, "+"))


	//字符串修改 
	s12 := "白萝卜"
	s13 := []rune(s12) //把字符串强制转换成一个rune切片
	s13[0] = '红'
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T \n",c1,c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n",c3,c4)
	fmt.Printf("%d\n",c4)
}

//字符串 ，必须双引号
// 字符，单个的字母汉子，字母。单引号