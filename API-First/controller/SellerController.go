package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/db"
	"github.com/zakverse/API-First/model"
)
func CreateSeller (c *gin.Context){
	var u model.Seller

	if err := c.ShouldBindJSON(&u);
	err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	result,err := db.DB.Exec("insert into seller(seller_code , nama , username)Values(?,?,?)", u.SellerCode , u.Nama , u.Username)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
    "message": "Success",
    "data":    result,
})
}

