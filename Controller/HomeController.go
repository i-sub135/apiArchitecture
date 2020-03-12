package controller

import (
	"net/http"

	r "../Repo"
	"github.com/gin-gonic/gin"
)

// RouteIndex -- index
func RouteIndex(res *gin.Context) {
	id := res.Param("id")
	res.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"result": map[string]string{
			"pesan": "Micro Services",
			"id":    id,
		},
	})
}

// RouteGetUser -- get user
func RouteGetUser(res *gin.Context) {
	data := r.GetUsers()
	res.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"result":  data,
	})
}

// RouteJwtUser -- get data jwt
func RouteJwtUser(res *gin.Context) {
	data := r.GetJwtUser()
	res.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "OK",
		"result":  data,
	})
}
