package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"message/src/cn/cncommdata/util/RandString"
)

/**
 * @author:libing.niu
 * @date: 2020/2/13 17:20
 */
type UserInfo struct {
	gorm.Model
	Name   string
	Gender string
	Hobby  string
}

//常量2
const two = 2

func main() {

	//连接数据库
	db, err := gorm.Open("mysql", "root:123456@/imcs?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("db link error :%v\n ", err.Error())
		return
	}
	fmt.Println(">>>数据库链接成功<<<")
	//删除数据库中的user_info表
	fmt.Println(">>>删除user_infos表<<<")
	db.DropTable("user_infos")
	// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
	fmt.Println(">>>创建user_infos表<<<")
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&UserInfo{})

	defer db.Close()
	//自动钱迁移
	AutoMigrate(db)
	//查询数据
	GetUserInfo(db)
}

//创建数据
func AutoMigrate(db *gorm.DB) {

	fmt.Println(">>>开始往user_infos表中写入数据...<<<")
	db.AutoMigrate(&UserInfo{})
	for i := 1; i <= 100; i++ {
		name := RandString.GetRandStrings(10)
		hobby := RandString.GetRandStrings(5)
		var gender string
		if i%two == 0 {
			//偶数设置为男
			gender = "男"
		} else {
			//奇数设置为女
			gender = "女"
		}
		u1 := UserInfo{Name: name, Gender: gender, Hobby: hobby}
		fmt.Printf("新增的第%v条数据为%v...\n", i, u1)
		//创建数据
		db.Create(&u1)
	}
	fmt.Println()

}

//查询数据
func GetUserInfo(db *gorm.DB) {
	fmt.Println(">>>开始从user_infos表中查询数据...<<<")
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("查询到的第一条数据为：%v\n", u)
	u = new(UserInfo)
	db.Find(u)
	fmt.Printf("查询到的最后一条数据为：%v\n", u)

}
