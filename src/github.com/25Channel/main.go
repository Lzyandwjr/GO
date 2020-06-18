package main

import "fmt"

// channel 是一种类型，一种引用类型。声明方式为： var 变量 chan 元素类型

// channel 可以在多个goroutine间进行通信，goroutine和channel是Go语言秉承CSP(Communicating Sequential Process)并发模式的重要实现基础

// Go语言的并发模型提倡通过通信共享内存而不是通过共享内存而实现的

// channel有发送（send）、接收(receive）和关闭（close）三种操作。 发送和接收都使用<-符号。

func recv(c chan int){
	ret := <- c
	fmt.Println("接收成功！",ret)
}

func main(){
	//var ch1 chan int // 通道是一种引用类型，需要初始化
	ch2 := make(chan int) // 无缓冲区通道,同步通道,如果直接获取数据，则会阻塞，则使用goroutine接受
	go recv(ch2)
	//ch1  =  make(chan int,1)// 有缓冲区通道
	//ch1 <- 10
	//x := <- ch1
	ch2 <- 200
	fmt.Println("发送成功")
	close(ch2)


}