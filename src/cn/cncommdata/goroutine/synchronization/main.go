package main

import (
	"fmt"
	"sync"
	"time"
)

/**

 * 协程同步 系统工具 sync.waitgroup
 *      1.Add(delta int) 添加协程记录
 *      2.Done() 移除协程记录
 *      3.Wait()同步等待所有记录的协程全部结束
 *
 * @author:libing.niu
 * @date: 2020/2/18 11:19
 */

var WG sync.WaitGroup

//读取
func Read() {

	for i := 1; i <= 3; i++ {
		fmt.Printf("添加第%v条协程\n", i)
		WG.Add(1)
	}

}

//写入
func Write() {
	for i := 1; i <= 3; i++ {
		time.Sleep(time.Second * 2)
		fmt.Println("Done--->", i)
		WG.Done()
		fmt.Printf("移除第%v条协程\n", i)

	}

}
func main() {

	//主线程读取文件
	Read()
	//启动协程写入文件
	go Write()
	//等待写入协程全部完成，继续向下执行
	WG.Wait()
	fmt.Println("all done")

}
