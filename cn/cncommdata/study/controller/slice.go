package controller

import (
	`container/list`
	`demo/cn/cncommdata/study/model`
	"fmt"
	`sync`
	`time`
)

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
func OriginSlice() {

	var numbers []int

	for i := 0; i < 20; i++ {
		numbers = append(numbers, append([]int{7, 8, 9}, []int{i}...)...)
		fmt.Println(numbers)
		fmt.Printf("切片长度%d，切片容量%d，切片地址值%p\r\n", len(numbers), cap(numbers), numbers)
	}

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

func PractiseMap() {

	strings := make(map[string][]model.EmailParam, 30)
	strings["name"] = []model.EmailParam{
		{ServerHost: "libing_niu@163.com"},
		{ServerPort: 645},
		{FromEmail: "43"},
		{FromPasswd: "4324"},
		{Toers: "ewr"},
		{CCers: "rerw"},
	}
	delete(strings, "name")

	fmt.Println(strings)
}

func PractiseSynicMap() {

	var scene sync.Map
	//添加元素
	scene.Store("name", "niulibing")
	scene.Store("sex", "男")
	fmt.Printf("map集合添加的元素为：%v\r\n", scene)

	//获取元素
	value, _ := scene.Load("name")

	fmt.Printf("通过name键获取到的元素为%v\r\n", value)

	//删除元素
	scene.Delete("sex")
	fmt.Println(scene)

}
func PractiseList() {

	//list 的初始化有两种方法：分别是使用 New() 函数和 var 关键字声明，两种方法的初始化效果都是一致的。
	//双链表支持从队列前方或后方插入元素，分别对应的方法是 PushFront 和 PushBack。

	i := list.New()

	i.PushBack(1)
	i.PushBack(2)

	fmt.Printf("golist插入的数据为%v", &i)

	//make和new的区别
	// make只能初始化go语言中的基本类型  make用于创建切片、哈希表、和管道等内置数据结构
	// new只接受一个类型作为参数，然后返回一个指向这个类型的指针  new用于分配并创建一个指向对应类型的指针

}

//可变参数
func Myfunc(a ...interface{}) {

	now := time.Now()

	for k, v := range a {
		fmt.Println(k, v)
	}

	since := time.Since(now)

	fmt.Printf("程序执行了%v", since)

}

type Wheel struct {
	Size int
}

type Engine struct {
	Power int
	Type  string
}

type Car struct {
	Wheel
	Engine
}

func StartCar() {

	car := new(Car)
	{

		// 初始化轮子
		Wheel{
			Size: 18,
		},
		// 初始化引擎
			Engine{
				Type:  "1.4T",
				Power: 143,
			},


			fmt.Print(car)
	}
}