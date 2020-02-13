package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {

	engine := gin.Default()

	engine.POST("/test", func(context *gin.Context) {

		all, err := ioutil.ReadAll(context.Request.Body)

		if err != nil {
			context.String(http.StatusBadRequest, err.Error())
			context.Abort()
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(all))
		firstName := context.PostForm("first_name")
		lastName := context.DefaultPostForm("last_name", "lirong")

		context.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(all))

	})

	engine.Run(":8086")
}
