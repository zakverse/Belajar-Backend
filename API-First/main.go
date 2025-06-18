package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/controller"
	"github.com/zakverse/API-First/db"
)
func main(){
    db.ConnectDatabase()

    r := gin.Default()
	r.POST("/seller",controller.CreateSeller)
    r.Run(":8080")
}