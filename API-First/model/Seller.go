package model

type Seller struct {
    ID         int    `json:"id"`           // opsional kalau auto-increment
    SellerCode string `json:"seller_code"`
    Nama       string `json:"nama"`
    Username   string `json:"username"`
}

