package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

var DB = make(map[string]string)

func ping(c *gin.Context) {
	fmt.Printf("Hello")
	c.String(200, "pong")
}

func user(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := DB[user]
	if ok {
		c.JSON(200, gin.H{"user": user, "value": value})
	} else {
		c.JSON(200, gin.H{"user": user, "status": "no value"})
	}
}

func admin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)

	var json struct {
		Value string `json:"value" binding:"required"`
	}

	if c.Bind(&json) == nil {
		DB[user] = json.Value
		c.JSON(200, gin.H{"status": "ok"})
	}

}