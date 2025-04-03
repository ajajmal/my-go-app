package main

import (
"github.com/gin-gonic/gin"
"log"
"net/http"
"github.com/ajajmal/my-go-app/config"
)

func main() {
	r:=gin.Default();
	r.GET("health", func(c *gin.Context){
		c.JSON(http.statusOK, gin.H{"status":"OK"})
	})

	if err := r.Run(":8080"); err != nil {
	log.Fatal(err)
	}

	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.User{})
}