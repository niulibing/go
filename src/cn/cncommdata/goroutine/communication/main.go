package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

/**
 * 日志监控系统「包括日志的读取、解析、写入三个模块，同时工作」
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
		//将读取到的数据发送至rc通道，实现读取模块和解析模块之间的数据通信

		rc <- readString[:len(readString)-1]
		fmt.Printf("数据读取模块从日志文件中读取到的数据:【%v】发送到rc通道。\n", readString)
	}
}

//实现写入接口
func (w *WriteToFluxDB) Write(wc chan string) {

	//写入模块中从通道wc中获取数据，进行数据处理。在此只做了数据打印
	fmt.Println("数据解析模块开始从wc通道中获取数据")
	for s := range wc {
		fmt.Printf("数据解析模块从wc通道中获取数据为:%v\n", s)
	}
}
func (l *LogProcess) Process() {

	//解析模块，从rc通道中解析数据，并且将解析出来的数据发送至wc通道与写入模块进行数据通信
	fmt.Println("数据解析模块开始从rc通道中获取数据")
	for s := range l.rc {
		fmt.Printf("数据解析模块从rc通道中获取数据为:%v\n", s)
		upper := strings.ToUpper(s)
		fmt.Printf("数据解析模块对rc通道中获取数据进行处理，处理后的数据为:%v并且发送到wc通道\n", upper)
		l.wc <- upper
	}
}
func main() {

	//查看该电脑cpu核数
	fmt.Printf("该电脑cpu为【%d】核...\n", runtime.NumCPU())
	//设置cpu的调度个数，如果不设置，默认调用全部cpu核数。
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	r := &ReadFromFile{path: "/Users/niulibing/Desktop/abc.log"}

	w := &WriteToFluxDB{influxDBDsn: "username&password"}

	lp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan string),
		read:  r,
		write: w,
	}
	//启动三个协程
	go lp.read.Read(lp.rc)
	fmt.Println("读取协程启动完成")
	go lp.Process()
	fmt.Println("解析协程启动完成")
	go lp.write.Write(lp.wc)
	fmt.Println("写入协程启动完成")
	//主程序延时关闭
	time.Sleep(120 * time.Second)

}
