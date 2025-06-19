package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/zakverse/API-First/db"
	"github.com/zakverse/API-First/model"
)
func CreateJournal (c *gin.Context){
	var u model.Journal

	if err := c.ShouldBindJSON(&u);
	err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	result,err := db.DB.Exec("insert into journal(seller_code , product_code , total , date)Values(?,?,?,?)", u.SellerCode, u.ProductCode , u.Total , u.Date)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
    "message": "Success",
    "data":    result,
})
}

func UpdateJournal (c *gin.Context){

	id := c.Param("id")
	var u model.Journal

	if err := c.ShouldBindJSON(&u);

	err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"massage" : err.Error()})
		return
	}
	_,err := db.DB.Exec("update seller set ... = ? , ....= ? where id = ? ",u.ProductCode, id , u.SellerCode , u.Total)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"massage":err.Error()})
		return
	}
	c.JSON(http.StatusOK , gin.H{
		"massage" : "Update Success",
	})


}

