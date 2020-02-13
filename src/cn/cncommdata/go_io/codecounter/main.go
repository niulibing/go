package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
 * @author:libing.niu
 * @date: 2020/2/13 17:18
 */
func main() {

	/*	fmt.Println(os.Args)
		//参数校验
		if len(os.Args) < 2 {
			fmt.Println("非法参数")
			return
		}
		fileName := os.Args[1]*/
	//打开文件
	file, err := os.Open("src/cn/cncommdata/go_io/codecounter/main1.go")

	if err != nil {
		fmt.Println(err)
		return
	}
	fileName := file.Name()
	//延时到整个main函数结束调用关闭
	defer file.Close()
	//文件读写器，读取文件中的每一行数据
	reader := bufio.NewReader(file)

	//循环读取文件中的每一行
	var line int
	for {
		_, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		if !prefix {

			line++
		}
	}

	fmt.Printf("文件%v，总共%v行", fileName, line)
}
