package model

type OperationParams struct {
	NewOwnerNick string `json:"NewOwnerNick"`
	CatID int `json:"CatID"`
	CatName string `json:"CatName"`
	PurchaseDate string `json:"PurchaseDate"`
	Status string `json:"Status"`
}
