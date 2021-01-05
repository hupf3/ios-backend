package models

import (
	"errors"
	"fmt"
)

// Bill 账单
type Bill struct {
	BillID   int     `json:"bill_id"`
	UserID   int     `json:"user_id"`
	Money    float32 `json:"money"`
	BillTime string  `json:"bill_time"`
	Classify string  `json:"classify"`
}

// InsertBill 插入账单信息
func InsertBill(b Bill) error {
	_, err := db.Exec("INSERT INTO bill(bill_id, user_id, money, classify) values(?, ?, ?, ?)", b.BillID, b.UserID, b.Money, b.Classify)

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

// GetAllBills 获取全部账单
func GetAllBills(userID int) ([]Bill, error) {
	bills := make([]Bill, 0)
	rows, err := db.Query("SELECT * FROM bill where user_id = ?", userID)
	if err != nil {
		fmt.Printf("Query bills failed, err:%v", err)
		return nil, err
	}
	for rows.Next() {
		var bill Bill
		if err = rows.Scan(&bill.BillID, &bill.UserID, &bill.Money, &bill.BillTime, &bill.Classify); err != nil {
			fmt.Printf("Scan bill failed, err:%v", err)
			return nil, err
		}
		bills = append(bills, bill)
	}
	return bills, nil
}

// QueryBill 获取账单信息
func QueryBill(billID int) (*Bill, error) {
	b := new(Bill)
	row := db.QueryRow("SELECT * FROM bill where bill_id = ?", billID)
	err := row.Scan(&b.BillID, &b.UserID, &b.Money, &b.Classify, &b.BillTime)

	if err != nil {
		fmt.Printf("Query bill failed, err:%v", err)
		return nil, errors.New("Bill does not exists")
	}

	return b, nil
}

// UpdateBill 修改账单信息
func UpdateBill(billID int, money float32, classify string) error {
	_, err := db.Exec("UPDATE bill SET money = ?, classify = ? WHERE bill_id = ?", money, classify, billID)

	if err != nil {
		fmt.Printf("Update bill failed, err:%v", err)
		return errors.New("Bill does not exists")
	}

	return nil
}
