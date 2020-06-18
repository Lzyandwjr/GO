package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client demo

func main(){
	// 与服务端简历连接
	conn,err := net.Dial("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	//连接成功，进行数据发送和接受
	input := bufio.NewReader(os.Stdin)
	for {
		s,_:= input.ReadString('\n')
		s = strings.TrimSpace(s)
		if strings.ToUpper(s) == "Q" {
			return
		}
		//给服务端发消息
		_,err := conn.Write([]byte(s))
		if err != nil {
			println(err)
			return
		}
		//从服务端接受回复的消息
		var buf [1024]byte
		n,err:=conn.Read(buf[:])
		if err !=nil{
			fmt.Println(err)
			return
		}
		fmt.Println("收到服务端回复:",string(buf[0:n]))


	}
}
