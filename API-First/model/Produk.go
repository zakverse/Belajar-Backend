package model

type Product struct {
	ProductCode 	string 		`json:"product_code"`
	ProductName 	string 		`json:"product_name"`
	Stock       	int    		`json:"stock"`
	Price       	int    		`json:"price"`
	CategoryCode 	string 		`json:"category_code"`
}