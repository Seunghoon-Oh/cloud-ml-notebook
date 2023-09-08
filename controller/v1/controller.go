package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-notebook-manager/service"
)

func GetNotebooks(c *gin.Context) {
	data := service.GetNotebooks()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func CreateNotebook(c *gin.Context) {
	data := service.CreateNotebook()
	c.String(http.StatusOK, data)
}
