package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/db"
	"github.com/zakverse/API-First/model"
)
func CreateProduct (c *gin.Context){
	var u model.Product

	if err := c.ShouldBindJSON(&u);
	err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	result,err := db.DB.Exec("insert into product(product_code , product_name , stock , price , category_code )Values(?,?,?,?,?)",
	 u.ProductCode , u.ProductName , u.Stock, u.Price , u.CategoryCode)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
    "message": "Success",
    "data":    result,
})
}

func UpdateProduct (c *gin.Context){

	id := c.Param("id")
	var u model.Product

	if err := c.ShouldBindJSON(&u);

	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage" : err.Error()})
		return
	}
	_,err := db.DB.Exec("update seller set ... = ? , ....= ? where id = ? ",u.ProductCode, u.ProductName, id ,u.Price , u.Stock)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	c.JSON(http.StatusOK , gin.H{
		"massage" : "Update Success",
	})


}

