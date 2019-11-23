package main

import (
	"demo/cn/cncommdata/study/stack"
	"fmt"
)

func main() {

	var arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ints := arr[:]
	arr = append(ints[0:2], ints[4:]...)
	fmt.Println(arr)

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
	stac()

}
func stac() {
	createStack := stack.CreateStack()
	createStack.Push(1313)
	createStack.Push(3443)
	fmt.Println(createStack.Pop(), createStack.Pop())
}
