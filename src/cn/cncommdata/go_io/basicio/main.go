package main

import (
	"fmt"
	"io"
	"strings"
)

func ReadFrom(reader io.Reader, num int) ([]byte, error) {

	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

//字符串读取的简单的例子
func sampleReadFromString() ([]byte, error) {

	from, er := ReadFrom(strings.NewReader("from string"), 12)
	return from, er

}

func main() {

	fromString, _ := sampleReadFromString()
	fmt.Println(fromString)
}
