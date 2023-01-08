package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/shipments", GetHandler)
	r.POST("/shipments", PostHandler)

	r.Run()
}
