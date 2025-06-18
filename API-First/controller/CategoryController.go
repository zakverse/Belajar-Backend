package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/db"
	"github.com/zakverse/API-First/model"
)

func CreateCategory(c *gin.Context) {
	var u model.Category

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	result, err := db.DB.Exec("insert into category_product(category_code, category_name)Values(?,?)", u.Category_code, u.Category_name)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Success",
		"data":    result,
	})
}
