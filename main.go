package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"msg": "Helloo"})
	})

	r.POST("/users/new", CreateUser)

	r.Run(":8080");
}
