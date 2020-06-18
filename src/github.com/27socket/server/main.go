package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp server demo

func process(con net.Conn){
	defer con.Close()//处理完要关闭连接
	//针对当前的连接做数据发送和接受
	for {
		reader := bufio.NewReader(con)
		var buf [128]byte
		n,err := reader.Read(buf[:])
		if err!=nil{
			fmt.Println(err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到客户端数据：",recv)
		con.Write([]byte("ok"))//把收到的数据返回给客户端
	}
}

func main(){
	listen,err := net.Listen("tcp","127.0.0.1:20000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		con,err := listen.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		//3.启动一个单独的goroutine去处理连接
		go process(con)
	}

}