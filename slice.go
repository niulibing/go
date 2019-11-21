package main

import (
	"fmt"
)

func main() {

	//array()
	//mySlice()
	//helloWorld()
	//originSlice()
	//resetSlice()
	//directStatementSlice()
	//useMakeFunConstructSlice()
	Map()

}
func helloWorld() {

	a := 1
	c := &a
	b := "hello world"

	fmt.Printf("变量a地址地址为 %s \r\n", &a)
	fmt.Printf("变量a值为 %v \r\n", *&a)
	fmt.Printf("变量b地址地址为 %s \r\n", &b)
	fmt.Printf("变量b值为 %v \r\n", *&b)
	fmt.Printf("=将a的值传递给c，c的值为：%v \r\n", *c)

	var arr [3]int

	arr[1] = 20
	arr[0] = 10
	arr[2] = 30
	fmt.Printf("数组的第一个元素的值为： %v \r\n", arr[0])
	fmt.Printf("数组的第二个元素的值为： %v \r\n", arr[1])
	fmt.Printf("数组的最后一个元素的值为： %v  \r\n", arr[len(arr)-1])

	//循环遍历数组
	for k, v := range arr {
		fmt.Print(k, v)
		fmt.Println()
	}

}
func array() {

	// 多维数组
	array1 := [4][2]int{{10, 11}, {20, 12}, {30, 134}, {404, 134}}

	fmt.Printf("该二维数组的长度为：%v \n", len(array1))
	for k, v := range array1 {

		fmt.Println(k, v)
	}

}

//切片
func mySlice() {

	//定义数组,
	var nameArr [30]int
	//给数组赋值
	for i := 0; i < 30; i++ {
		nameArr[i] = i + 1
	}
	i2 := append(nameArr[:], 1, 2, 3)
	fmt.Printf("nihao%v\r\n", i2)
	//遍历数组

	//切片
	ints := nameArr[0:2]
	fmt.Printf("数组的值为 %v\r\n切片的值为%v\r\n", nameArr, ints)
	//将切片再次切片
	i := ints[0:1]
	fmt.Printf("ints切片再次切片的结果为%v\r\n", i)
	//区间
	fmt.Printf("区间切片%v\r\n", nameArr[10:15])
	//中间到尾部的所有元素
	fmt.Printf("中间到尾部的所有元素%v\r\n", nameArr[20:])
	//开头到中间指定位置的元素
	fmt.Printf("开头到中间到指定位置的所有元素%v\r\n", nameArr[:10])
}

//原始切片
func originSlice() {

	a := []int{1, 2, 3}
	fmt.Printf("原始切片为%v", a[:])
}

//重置切片
func resetSlice() {
	a := []int{1, 2, 3}
	fmt.Printf("原始切片为%v，重置后的切片为%v", a[:], a[0:0])
}

//直接声明切片
func directStatementSlice() {

	//声明字符串切片
	var strList []string
	//声明整形切片
	var numberList []int
	//声明一个空切片   分配了内存，所以不是空的
	var numberListEmpty = []int{}
	//输出3个切片
	fmt.Println(strList, numberList, numberListEmpty)

	//输出3个切片的大小

	fmt.Println(len(strList), len(numberList), len(numberListEmpty))

	//切片判定空的结果

	fmt.Println(strList == nil)
	fmt.Println(numberList == nil)
	fmt.Println(numberListEmpty == nil)

	//获取内存地址

	fmt.Printf("strList的内存地址为%p；numberList的内存地址为%p;numberListEmpty的内存地址为%p\n", &strList, &numberList, &numberListEmpty)
	fmt.Printf("strList的内存地址为%p；numberList的内存地址为%p;numberListEmpty的内存地址为%p", strList, numberList, numberListEmpty)
}

//使用make函数构造切片
func useMakeFunConstructSlice() {

	//如果需要动态地创建一个切片，可以使用make函数
	//其中a和b均是预分配2个元素的切片，只是b的内部存储空间已经分配了10个，但实际是用来2个元素
	b := make([]int, 2, 10)
	a := make([]int, 2)
	fmt.Println(len(b), len(a))
	fmt.Println(b, a)

	//温馨提示： 使用make函数生成的切片一定发生了内存分配，但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作
}

//使用append为切片添加元素
func useAppendAddElement() {

}
