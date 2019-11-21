package main

import "fmt"

func main() {

	//a := 1
	//c := &a
	//b := "hello world"
	//
	//fmt.Printf("变量a地址地址为 %s \r\n", &a)
	//fmt.Printf("变量a值为 %v \r\n", *&a)
	//fmt.Printf("变量b地址地址为 %s \r\n", &b)
	//fmt.Printf("变量b值为 %v \r\n", *&b)
	//fmt.Printf("=将a的值传递给c，c的值为：%v \r\n", *c)
	//
	//var arr [3]int
	//
	//arr[1] = 20
	//arr[0] = 10
	//arr[2] = 30
	//fmt.Printf("数组的第一个元素的值为： %v \r\n", arr[0])
	//fmt.Printf("数组的第二个元素的值为： %v \r\n", arr[1])
	//fmt.Printf("数组的最后一个元素的值为： %v  \r\n", arr[len(arr)-1])
	//
	////循环遍历数组
	//for k, v := range arr {
	//	fmt.Print(k, v)
	//	fmt.Println()
	//}
	//array()

	send()

}

func array() {

	// 多维数组
	array1 := [4][2]int{{10, 11}, {20, 12}, {30, 134}, {404, 134}}

	fmt.Printf("该二维数组的长度为：%v \n", len(array1))
	for k, v := range array1 {

		fmt.Println(k, v)
	}

}
