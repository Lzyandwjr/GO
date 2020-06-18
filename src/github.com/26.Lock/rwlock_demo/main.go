package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var(
	x int64
	rwlock sync.RWMutex
	wg sync.WaitGroup
	lock sync.Mutex
)

func read(){
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()
	wg.Done()
}

func write(){
	rwlock.RLock()
	x = x+1
	time.Sleep(time.Millisecond*10)
	rwlock.RUnlock()
	wg.Done()
}

func main(){
	start := time.Now()
	for i := 0;i<1000;i++{
		wg.Add(1)
		go read()
	}

	for i := 0;i<10;i++{
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}