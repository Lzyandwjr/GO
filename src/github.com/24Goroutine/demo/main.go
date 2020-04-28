package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup

func main(){	// 开启一个主gorountine去执行main函数

	wg.Add(100000) //计数器+1
	for i:=0;i<=100000 ;i++  {
		go func(i int) {
			fmt.Println("hello goruntine!",i)
			wg.Done()//计数器-1
		}(i)
	}
	// 开启了一个goroutine去执行hello函数
	fmt.Println("heasdfasfasddfsadfasfasdfasdllo")
	//time.Sleep(time.Second)//有可能会主函数已经运行完了，而另外一个线程并没有完成的现象
	wg.Wait()//阻塞，等所有线程干完才结束

}
