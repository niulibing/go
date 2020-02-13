package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()
	engine.GET("/get", func(c *gin.Context) {
		c.String(200, "this is get request")
	})
	engine.POST("/post", func(c *gin.Context) {
		c.String(200, "this is post request")
	})
	engine.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(200, "this is delete request")
	})

	engine.Any("/any", func(context *gin.Context) {
		context.String(200, "this is any request")
	})
	engine.Run(":8081")
}
