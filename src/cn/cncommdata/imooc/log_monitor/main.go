package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

/**
 * 日志监控系统
 * @author:libing.niu
 * @date: 2020/2/17 14:47
 */

//声明读接口
type Reader interface {
	Read(rc chan string)
}

//声明写接口
type Writer interface {
	Write(wc chan string)
}

//声明从文件读取的结构体
type ReadFromFile struct {
	path string
}
type LogProcess struct {
	rc    chan string //读取察南
	wc    chan string
	read  Reader
	write Writer
}

//声明解析文件的结构体
type WriteToFluxDB struct {
	//influx dbata source
	influxDBDsn string
}

//实现读取接口
func (r *ReadFromFile) Read(rc chan string) {

	//打开日志文件
	file, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error: %s", err.Error()))
	}
	//将追针移动到文件末尾
	file.Seek(0, 2)
	reader := bufio.NewReader(file)
	//从文件末尾逐行读取
	for {
		readString, err := reader.ReadString('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("ReadString error:%s", err.Error()))
		}
		rc <- readString
	}
}

//实现写入接口
func (w *WriteToFluxDB) Write(wc chan string) {

	for s := range wc {

		fmt.Println(s)
	}
}
func (l *LogProcess) Process() {
	//解析模块

	for s := range l.rc {
		l.wc <- s
	}
}
func main() {

	r := &ReadFromFile{path: "/Users/niulibing/Desktop/abc.log"}

	w := &WriteToFluxDB{influxDBDsn: "username&password"}

	lp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(120 * time.Second)

}
