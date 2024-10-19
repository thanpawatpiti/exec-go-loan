package models

import (
	"errors"

	"github.com/thanpawatpiti/exec-go-loan/pkg/handlers/usecase"
	orm "github.com/thanpawatpiti/exec-go-loan/pkg/models/ORM"
	"golang.org/x/crypto/bcrypt"
)

func (m *Init) Login(body usecase.RequestLogin) (*uint, error) {
	var accountUser orm.AccountUser

	// ค้นหา user จาก username (account_value)
	err := m.DB.Model(&orm.AccountUser{}).Where("account_value = ?", body.Username).First(&accountUser).Error
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	if accountUser.AccountType != "username" {
		return &accountUser.UserID, nil
	}

	// ตรวจสอบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(accountUser.EncryptedPassword), []byte(body.Password))
	if err != nil {
		return nil, errors.New("invalid username or password")
	}

	return &accountUser.UserID, nil
}
