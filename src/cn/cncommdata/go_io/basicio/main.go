package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/**
 * @author:libing.niu
 * @date: 2020/2/13 17:16
 */

func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

//字符串读取的简单的例子
func sampleReadFromString() {
	fromString, _ := ReadFrom(strings.NewReader("from string"), 12)
	fmt.Println(string(fromString))

}

//读取从命令行的输入的数据
func sampleReadFromStdin() {
	fmt.Println("please input from stdin:")
	stdin, _ := ReadFrom(os.Stdin, 11)
	fmt.Println(string(stdin))
}

//从文件读取
func sampleReadFile() {

	file, err := os.Open("src/cn/cncommdata/go_io/basicio/main1.go")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fileContext, _ := ReadFrom(file, 50)
	fmt.Println(string(fileContext))

}

func main() {

	sampleReadFromString()
	sampleReadFromStdin()
	sampleReadFile()
}
