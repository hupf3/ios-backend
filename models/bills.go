package models

// Bill 账单
type Bill struct {
	BillID   int    `json:"bill_id"`
	UserID   int    `json:"user_id"`
	Money    string `json:"money"`
	BillTime string `json:"bill_time"`
	Tag      string `json:"tag"`
}
