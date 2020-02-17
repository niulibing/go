package main

import "fmt"

/**
 * @author:libing.niu
 * @date: 2020/2/14 11:36
 */
type School struct {
	name string
}
type Student struct {
	School
	name string
	age  int
}

func (stu *Student) register() {
	fmt.Printf("该学生隶属于：%v,姓名：%v，年龄为：%v", stu.School.name, stu.name, stu.age)
}
func main() {

	student := Student{School{name: "北京大学"}, "zhansan", 20}
	student.register()
}
