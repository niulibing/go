package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	engine := gin.Default()
	engine.GET("/test", func(context *gin.Context) {

		// get params
		firstName := context.Query("first_name")
		lastName := context.DefaultQuery("last_name", "libing")
		//output params
		context.String(http.StatusOK, "第一个名称为：%s，第二个名称为：%s", firstName, lastName)

	})

	//start server
	engine.Run(":8085")
}
