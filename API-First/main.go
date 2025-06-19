package main

import (
	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/controller"
	"github.com/zakverse/API-First/db"
)
func main(){
    db.ConnectDatabase()

    r := gin.Default()
	r.POST("/seller",controller.CreateSeller) //Post For Create (Insert)
	r.PATCH("/seller/:id",controller.UpdateSeller) // Patch For Update

	r.POST("/produk",controller.CreateProduct)
	r.PATCH("/produk/:id",controller.UpdateProduct)


	r.POST("/category",controller.CreateCategory)
	r.PATCH("/category/:id",controller.UpdateCategory)


	r.POST("/journal",controller.CreateJournal)
	r.PATCH("/journal/:id",controller.UpdateJournal)


    r.Run(":8080")
}