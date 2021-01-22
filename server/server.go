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

func InitServer() {
	router := gin.Default()
	//v1组路由
	v1 := router.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/submit", submitEndpoint)
		v1.GET("/read", readEndpoint)
	}

	//v2组路由
	v2 := router.Group("/v2")
	{
		v2.GET("/login", loginEndpoint)
		v2.GET("/submit", submitEndpoint)
		v2.GET("/read", readEndpoint)
	}
	router.Run()
}
