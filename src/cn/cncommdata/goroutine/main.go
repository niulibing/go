package main

import (
	"fmt"
	"runtime"
)

/**
 * 并发版本
 * @author:libing.niu
 * @date: 2020/2/17 16:19
 */

func main() {

	fmt.Printf("cpu num = %d", runtime.NumCPU())
	//设置cpu的最大调度个数。如果不设置，cpu是满载调度(还是要给系统留出部分资源)
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	//初始化chan
	/*	ch := make(chan string)
		for i := 0; i < 5; i++ {
			go PrintHelloWorld(i, ch)
		}
		//从通道中输出消息
		for {
			fmt.Println(<-ch)
		}*/

}

func PrintHelloWorld(i int, ch chan string) {

	for {
		ch <- fmt.Sprintf("hello world from goroutine: %v", i)
	}

}
