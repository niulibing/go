package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()

	//绑定文件夹夹的第一种写法
	engine.Static("/assets", "./assets")
	//绑定文件夹的第二种写法
	engine.StaticFS("/static", http.Dir("static"))

	engine.Run(":8083")
}
