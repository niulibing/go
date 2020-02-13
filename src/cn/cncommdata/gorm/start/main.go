package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
 * @author:libing.niu
 * @date: 2020/2/13 17:20
 */
type Products struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {

	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/imcs?charset=utf8&parseTime=True&loc=Local")
	db.DropTable("products")
	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Products{})
	if err != nil {

		panic("连接数据库失败")
	}
	fmt.Println("数据库链接成功")
	//defer db.Close()

}
