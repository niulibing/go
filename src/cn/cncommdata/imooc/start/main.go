package main

import (
	"fmt"
	"strings"
	"time"
)

/**
 * @author:libing.niu
 * @date: 2020/2/14 14:52
 */

type LogProcess struct {
	rc          chan string //读取察南
	wc          chan string
	path        string //读取文件路径
	influxDBDsn string //influx dbata source

}

func (l *LogProcess) ReadFromFile() {

	//读取模块
	line := "message"
	//将消息放至rcchanel中将消息传递
	l.rc <- line

}
func (l *LogProcess) Process() {
	//解析模块
	data := <-l.rc
	l.wc <- strings.ToUpper(data)
}
func (l *LogProcess) WriteToInfluxDB() {
	//写入模块

	fmt.Println(<-l.wc)
}
func main() {

	lp := &LogProcess{
		rc:          make(chan string),
		wc:          make(chan string),
		path:        "/temp/access.log",
		influxDBDsn: "username&password",
	}
	go lp.ReadFromFile()
	go lp.Process()
	go lp.WriteToInfluxDB()

	time.Sleep(1 * time.Second)

}
