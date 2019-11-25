package main

import (
	"fmt"
	"message/src/cn/cncommdata/study/controller"
	"message/src/cn/cncommdata/study/stack"
)

func main() {

	//array()
	//mySlice()
	//helloWorld()
	//originSlice()
	//resetSlice()
	//directStatementSlice()
	//useMakeFunConstructSlice()
	//utils.Send()

	//实例化file

	//file := model.FileConstruct(true)
	//
	//fmt.Println(file)
	//
	////声明一个DataWriter的接口
	//var writer service.DataWriter
	//
	////将接口赋值file，也就是*file类型
	//writer = file
	//
	//b := file.CanOrNo
	//if b {
	//	write := writer.CanWrite(b)
	//	if write {
	//		data := writer.WriteData("欢迎写入数据")
	//		fmt.Println(data)
	//	} else {
	//		fmt.Printf("不好意思，你无权写入")
	//	}
	//}

	//controller.OriginSlice()
	//fmt.Println("主程序执行")
	//stac()
	//controller.PractiseMap()
	controller.PractiseSynicMap()

	//controller.PractiseList()
	//controller.Myfunc(1,2,3,4,5,67,8,"hello go",make(map[int]int))
	//controller.StartCar()

}
func stac() {
	createStack := stack.CreateStack()
	createStack.Push(1313)
	createStack.Push(3443)
	fmt.Println(createStack.Pop(), createStack.Pop())
}
