package main

import (
	"github.com/gin-gonic/gin"

	Controller "./Controller"
)

func main() {
	route := gin.Default()
	route.GET("/", Controller.RouteIndex)

	route.Run()
}
