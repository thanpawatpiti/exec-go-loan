package models

import (
	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"
	orm "github.com/thanpawatpiti/exec-go-loan/pkg/models/ORM"
)

func (m *Init) GetProfile(userID uint) (*usecase.ResponseProfile, error) {
	var res usecase.ResponseProfile

	// ดึงข้อมูลผู้ใช้จากฐานข้อมูล
	var user orm.Users
	err := m.DB.Model(&orm.Users{}).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	// ตั้งค่าข้อมูลใน response
	res.FistName = user.FistName
	res.LastName = user.LastName

	// ดึง account ของผู้ใช้
	var accounts []orm.AccountUser
	err = m.DB.Model(&orm.AccountUser{}).Where("user_id = ?", userID).Find(&accounts).Error
	if err != nil {
		return nil, err
	}

	// ตั้งค่าข้อมูล account ใน response
	for _, account := range accounts {
		res.Accounts = append(res.Accounts, usecase.ResponseProfileAccount{
			AccountType:  account.AccountType,
			AccountValue: account.AccountValue,
		})
	}

	return &res, nil
}
