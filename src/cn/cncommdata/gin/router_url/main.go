package main

import "github.com/gin-gonic/gin"

func main() {

	engine := gin.Default()
	//restful request type
	engine.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})
	engine.Run(":8082")
}
