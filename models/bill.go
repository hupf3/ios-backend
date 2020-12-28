package models

import (
	"errors"
	"fmt"
)

// Bill 账单
type Bill struct {
	BillID   int    `json:"bill_id"`
	UserID   int    `json:"user_id"`
	Money    string `json:"money"`
	BillTime string `json:"bill_time"`
	Tag      string `json:"tag"`
}

// InsertBill 插入账单信息
func InsertBill(b Bill) error {
	_, err := db.Exec("insert INTO bill(bill_id, user_id, money, tag) values(?,?)", b.BillID, b.UserID, b.Money, b.Tag)

	if err != nil {
		fmt.Printf("Insert bill failed, err:%v", err)
		return errors.New("Bill exists")
	}

	return nil
}
