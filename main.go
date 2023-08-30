package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-notebook-manager/controller/v1"
	"github.com/gin-gonic/gin"
)

// func HomePage(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "홈페이지 입니다.",
// 	})
// }

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.String(200, "pong")
	// })

	r.GET("/notebooks", v1.GetNotebooks)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8082")
}
