package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", ping)
	r.GET("/user/:name", user)

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"frazerbayley":  "pw",
	}))

	authorized.POST("admin", admin)

	return r
}