package models

import (
	"errors"
	"fmt"
)

// Bill 账单
type Bill struct {
	BillID   int    `json:"bill_id"`
	UserID   int    `json:"user_id"`
	Money    int    `json:"money"`
	BillTime string `json:"bill_time"`
	Tag      string `json:"tag"`
}

// InsertBill 插入账单信息
func InsertBill(b Bill) error {
	_, err := db.Exec("INSERT INTO bill(bill_id, user_id, money, tag) values(?, ?, ?, ?)", b.BillID, b.UserID, b.Money, b.Tag)

	if err != nil {
		fmt.Printf("Insert bill failed, err:%v", err)
		return errors.New("Bill exists")
	}

	return nil
}

// DeleteBill 删除账单信息
func DeleteBill(billID int) error {
	_, err := db.Exec("DELETE FROM bill where bill_id = ?", billID)

	if err != nil {
		fmt.Printf("Delete bill failed, err:%v", err)
		return errors.New("Bill does not exists")
	}

	return nil
}

// QueryBill 获取账单信息
func QueryBill(billID int) (*Bill, error) {
	b := new(Bill)
	row := db.QueryRow("SELECT * FROM bill where bill_id = ?", billID)
	err := row.Scan(&b.BillID, &b.UserID, &b.Money, &b.Tag, &b.BillTime)

	// fmt.Println(&b.BillID)

	if err != nil {
		fmt.Printf("Query bill failed, err:%v", err)
		return nil, errors.New("Bill does not exists")
	}

	return b, nil
}

// UpdateBill 修改账单信息
func UpdateBill(billID int, money int, tag string) error {
	_, err := db.Exec("UPDATE bill SET money = ?, tag = ? WHERE bill_id = ?", money, tag, billID)

	if err != nil {
		fmt.Printf("Update bill failed, err:%v", err)
		return errors.New("Bill does not exists")
	}

	return nil
}
