package main

import "github.com/gin-gonic/gin"

func main() {

	//创建gin的实例
	engine := gin.Default()

	engine.GET("/ping", func(context *gin.Context) {

		context.JSON(200, gin.H{

			"message": "require sucess",
		})
	})

	engine.Run(":8080")
}
