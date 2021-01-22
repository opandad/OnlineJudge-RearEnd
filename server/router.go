package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func loginEndpoint(c *gin.Context) {
	fmt.Println("这是login方法")
}

func submitEndpoint(c *gin.Context) {
	fmt.Println("这是submit方法")
}

func readEndpoint(c *gin.Context) {
	fmt.Println("这是read方法")
}

func index() {
	fmt.Println("???")
}

func InitServer() {
	router := gin.Default()
	router.GET("/")
	router.GET("/login")
	router.GET("/regist")
	router.GET("/forget_password")

	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}
	router.Run()
}
