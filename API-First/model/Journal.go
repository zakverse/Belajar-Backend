package model

type Journal struct {
	SellerCode 		string 		`json:"seller_code"`
	ProductCode 	string 		`json:"product_code"`
	Total 			int 		`json:"total"`
	Date 			string 		`json:"date"`
}