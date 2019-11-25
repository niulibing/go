package model

import "fmt"

type File struct {
	CanOrNo bool
}

func FileConstruct(bo bool) *File {

	return &File{CanOrNo: bo}
}

func (f *File) CanWrite(bo interface{}) bool {

	if f.CanOrNo {
		fmt.Printf("允许数据写入。。。。")
		return true
	}
	fmt.Printf("不允许数据写入....")
	return false
}

func (f *File) WriteData(data interface{}) error {

	//模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}
