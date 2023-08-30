package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Seunghoon-Oh/cloud-ml-notebook-manager/service"
)

func GetNotebooks(c *gin.Context) {
	data := service.GetNotebooks()
	println("Response: " + data)
	c.String(http.StatusOK, data)
}
