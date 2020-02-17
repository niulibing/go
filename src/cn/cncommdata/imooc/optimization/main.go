package main

import (
	"fmt"
	"strings"
	"time"
)

/**
 * 并发编程代码优化
 * @author:libing.niu
 * @date: 2020/2/14 16:24
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
	line := r.path
	rc <- strings.ToUpper(line)
}

//实现写入接口
func (w *WriteToFluxDB) Write(wc chan string) {

	fmt.Println(<-wc)
}
func (l *LogProcess) Process() {
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}
func main() {

	r := &ReadFromFile{path: "/temp/access.log",
	}

	w := &WriteToFluxDB{influxDBDsn: "username&password",
	}

	lp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}
	go lp.read.Read(lp.rc)
	go lp.Process()
	go lp.write.Write(lp.wc)

	time.Sleep(1 * time.Second)

}
