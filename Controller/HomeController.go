package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteIndex -- index
func RouteIndex(res *gin.Context) {
	res.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"result": map[string]string{
			"pesan": "Micro Services",
		},
	})
}
