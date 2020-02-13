package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/**
 * @author:libing.niu
 * @date: 2020/2/13 17:17
 */
func bufiosample() {

	//字符串读取器
	strReader := strings.NewReader("hello world")
	bufReader := bufio.NewReader(strReader)
	peek, _ := bufReader.Peek(5)
	fmt.Println(string(peek), bufReader.Buffered())
	readString, _ := bufReader.ReadString(' ')
	fmt.Println(string(readString))

	//写出器
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(writer, "你好")
	fmt.Fprint(writer, "世界")
	writer.Flush()

}

func main() {
	bufiosample()
}
