package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	Controller "./Controller"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	route := gin.Default()

	route.GET("/", Controller.RouteIndex)
	route.GET("/users", Controller.RouteGetUser)
	route.GET("/jwt", Controller.RouteJwtUser)

	route.Run()
}
