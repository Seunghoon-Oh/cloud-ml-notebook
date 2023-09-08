package main

import (
	v1 "github.com/Seunghoon-Oh/cloud-ml-notebook-manager/controller/v1"
	"github.com/Seunghoon-Oh/cloud-ml-notebook-manager/data"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	r.GET("/notebooks", v1.GetNotebooks)
	r.POST("/notebook", v1.CreateNotebook)

	return r
}

func main() {
	data.SetupRedisClient()
	r := setupRouter()
	r.Run(":8082")
}
