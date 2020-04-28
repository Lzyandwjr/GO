package main

import (
	"fmt"
	"sync"
)

// go语言并发编程
/*
	并发：同一时间段内执行多个任务
    并行：同一时刻执行多个任务

	一个函数可以被创建多个goruntine，一个goroutine必定对应一个函数
*/
var wg sync.WaitGroup
func hello(i int){
	fmt.Println("hello goruntine!",i)
	wg.Done()//计数器-1
}

func main(){	// 开启一个主gorountine去执行main函数


	wg.Add(10000) //计数器+1
	for i:=0;i<=10000 ;i++  {
		go hello(i)
	}
		// 开启了一个goroutine去执行hello函数
	fmt.Println("heasdfasfasddfsadfasfasdfasdllo")
	//time.Sleep(time.Second)//有可能会主函数已经运行完了，而另外一个线程并没有完成的现象
	wg.Wait()//阻塞，等所有线程干完才结束

}

