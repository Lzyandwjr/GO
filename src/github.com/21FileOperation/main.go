package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件操作

func readFromFile(){
	fileobj,err := os.Open("./xx.txt")
	if err !=nil{
		fmt.Println("open file failed,err:%v\n",err)
		return
	}
	//关闭文件
	defer fileobj.Close()
	// 读取文件内容
		var tmp = make([]byte,128)
		n,err := fileobj.Read(tmp)
		if err != nil{
			fmt.Println("open file failed,err:%v\n",err)
			return
		}
		fmt.Printf("read %d bytes from file.\n",n)
		fmt.Println(string(tmp))

}
func readAll(){
	fileobj,err := os.Open("./xx.txt")
	if err !=nil{
		fmt.Println("open file failed,err:%v\n",err)
		return
	}
	//关闭文件
	defer fileobj.Close()
	// 读取文件内容
	for {
		var tmp = make([]byte,128)
		n,err := fileobj.Read(tmp)
		if err == io.EOF{// End of File
			//把当前读了多少个字节的数据打印出来人后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil{
			fmt.Println("open file failed,err:%v\n",err)
			return
		}
		fmt.Printf("read %d bytes from file.\n",n)
		fmt.Println(string(tmp[:n]))
	}
}

// read by bufio
func readBufio(){
	fileobj,err := os.Open("./xx.txt")
	if err !=nil{
		fmt.Println("open file failed,err:%v\n",err)
		return
	}
	//关闭文件
	defer fileobj.Close()
	reader := bufio.NewReader(fileobj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF{
			fmt.Print(line)
			return
		}
		if err != nil {
			fmt.Printf("read file by bufio failed,err:%v\n", err)
			return
		}
		fmt.Print(line)
	}
}
// ioutil
func readFromIoutil(){
	content,err := ioutil.ReadFile("./xx.txt")
	if err != nil{
		fmt.Println("出错")
		return
	}
	fmt.Println(string(content))
}
// 文件写入操作

func wirte(){
	fileObj,err := os.OpenFile("./xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)

	if err != nil{
		fmt.Println("wrong")
		return
	}
	defer fileObj.Close()
	str := "lzy oooo"
	fileObj.Write([]byte(str))
	fileObj.WriteString("hello world")//写入字符串

}

func writeByBufio(){
	fileObj,err := os.OpenFile("./xx.txt",os.O_CREATE|os.O_WRONLY|os.O_APPEND,0644)

	if err != nil{
		fmt.Println("wrong")
		return
	}
	defer fileObj.Close()
	writer := bufio.NewWriter(fileObj)
	writer.WriteString("hhahahahah")
	writer.Flush() // 将缓冲区的内容写入磁盘
}

// 无需打开或者关闭文件
func writeByIoutil(){
	str := "66666666666666666"
	err := ioutil.WriteFile("./xx.txt",[]byte(str),0644)
	if err != nil{
		return
	}
}

func main(){
	//readAll()
	//readBufio()
	//readFromIoutil()
	//wirte()
	//writeByBufio()
	writeByIoutil()
}
//func main(){
//	// 打开文件

//}
