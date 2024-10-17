package orm

import (
	"time"
)

type AccountUser struct {
	ID                uint      `gorm:"primaryKey;autoIncrement;column:id"`
	UserID            uint      `gorm:"column:user_id;not null"`
	AccountType       string    `gorm:"column:account_type;not null"`
	AccountValue      string    `gorm:"column:account_value;not null"`
	EncryptedPassword string    `gorm:"column:encrypted_password;not null"`
	CreatedAt         time.Time `gorm:"column:created_at;not null"`
	UpdatedAt         time.Time `gorm:"column:updated_at;not null"`
}

func (a *AccountUser) TableName() string {
	return "account_user"
}
