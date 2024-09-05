package dtos

type TransactionDetail struct {
	Quantity    int `json:"quantity" form:"quantity"`
	Product     int `json:"product"`
	Transaction int `json:"transaction" form:"transaction"`
	Variant     int `json:"variant" form:"variant"`
	ProductSize int `json:"productSize" form:"productSize"`
}

type FormTransaction struct {
	FullName          string `json:"fullName" form:"fullName"`
	Email             string `json:"email" form:"email"`
	Address           string `json:"address" form:"address"`
	Payment           string `json:"payment" form:"payment"`
	OrderType         int    `json:"orderType" form:"orderType"`
	TransactionStatus int    `json:"transactionStatus" form:"transactionStatus"`
}
