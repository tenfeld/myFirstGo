package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.GET("/user/:id", getUserData)

	router.Run(":8080")
}

func getUserData(c *gin.Context) {
	user_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(200, gin.H{"user_id":"0"})
		return
	}
	c.JSON(200, gin.H{"user_id": user_id})
}