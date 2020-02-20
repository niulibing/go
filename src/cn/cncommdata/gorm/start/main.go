package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"message/src/cn/cncommdata/util/RandString"
	"time"
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
	Age    int
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

	/*db.DropTable("user_infos")
	fmt.Println(">>>删除user_infos表<<<")*/

	//如果不存在user_infos表，则进行创建
	if !db.HasTable("user_infos") {
		// 创建表`users'时将“ENGINE = InnoDB”附加到SQL语句
		fmt.Println(">>>创建user_infos表<<<")
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&UserInfo{})
	}

	defer db.Close()
	//创建数据
	/*	//AutoMigrate(db)*/
	/*	//查询第一条数据*/
	/*	QueryFirst(db)*/
	/*	//查询最后一条数据*/
	/*	QueryLast(db)*/
	/*	//根据id查询*/
	/*	QueryById(db, 200)*/
	/*	//查询所有数据*/
	/*	QueryAll(db)*/
	/*	//简单条件查询*/
	/*	QueryBySimpleCondition(db)*/
	QueryByStructAndMap(db)
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
		u1 := UserInfo{Name: name, Gender: gender, Hobby: hobby, Age: i}
		fmt.Printf("新增的第%v条数据为%v...\n", i, u1)
		//创建数据
		db.Create(&u1)
	}
	fmt.Println()

}

//查询数据
func QueryFirst(db *gorm.DB) {
	fmt.Println(">>>开始从user_infos表中查询数据...<<<")
	// 查询
	var u = new(UserInfo)
	db.First(&u)
	fmt.Printf("第一条数据为id为【%v】：%v\n", u.ID, u)

}
func QueryById(db *gorm.DB, id int) {

	//根据id查询
	u := new(UserInfo)
	db.First(&u, id)
	fmt.Printf("根据id【%v】查询的数据为：%v\n", u.ID, u)
}
func QueryLast(db *gorm.DB) {

	//根据id查询
	u := new(UserInfo)
	db.Last(&u)
	fmt.Printf("最后一条数据为id为【%v】，数据为：%v\n", u.ID, u)
}

//查询所有数据
func QueryAll(db *gorm.DB) {

	var users []UserInfo
	db.Find(&users)
	for _, user := range users {
		fmt.Printf("数据为id为【%v/%v】，数据为：%v\n", user.ID, len(users), user)
	}

}

//简单条件查询
func QueryBySimpleCondition(db *gorm.DB) {

	//声明结构体切片
	var users []UserInfo
	//声明结构体
	var user UserInfo

	//1.获取满足条件的所有记录
	db.Where("name=?", "PVSXSSDKBJ").Find(&users)
	fmt.Printf("符合条件的记录的总条数为【%v】条，分别为%v\n", len(users), users)

	//2.获取第一条匹配记录
	db.Where("name=?", "PVSXSSDKBJ").First(&user)
	fmt.Printf("符合条件的第一条记录为：%v\n", user)

	//3.IN
	db.Where("name in (?)", []string{"PVSXSSDKBJ", "CORCBSEJGH"}).Find(&users)
	fmt.Printf("name在PVSXSSDKBJ,CORCBSEJGH有【%v】条，分别为%v\n", len(users), users)

	//4.LIKE

	db.Where("name LIKE ?", "%RR%").Find(&users)
	fmt.Printf("name Like RR 有【%v】条，分别为%v\n", len(users), users)

	//5.AND
	db.Where("name=? and age>=?", "OGFNYVZWVN", "98").Find(&users)
	fmt.Printf("name为OGFNYVZWVN并且年龄大于98岁的有【%v】条，分别为%v\n", len(users), users)

	//6.Time     time.Now()
	db.Where("updated_at>?", time.Date(2020, time.February, 20, 17, 00, 00, 999, time.UTC)).Find(&users)
	fmt.Printf("updated_at >time.Now()有【%v】条", len(users))
}

//结构体、map、切片查询
func QueryByStructAndMap(db *gorm.DB) {

	//当使用struct查询时，GORM将只查询那些具有值的字段

	var user UserInfo
	var users []UserInfo

	//struct
	db.Where(UserInfo{Name: "HWBIWQMVXP", Age: 85}).First(&user)
	fmt.Printf("结构体查询结果为：%v\n", user)
	//Map
	db.Where(map[string]interface{}{"name": "HWBIWQMVXP", "age": "85"}).Find(&users)
	fmt.Printf("map查询结果为【%v】条：%v\n", len(users), users)
	//主键的切片查询
	db.Where([]int64{1, 500, 1000}).Find(&users)
	fmt.Printf("主键切片查询结果为【%v】条：%v\n", len(users), users)

}
