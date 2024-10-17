package orm

import "time"

type Users struct {
	ID        uint      `gorm:"primaryKey;autoIncrement;column:id"`
	FistName  string    `gorm:"column:first_name;not null"`
	LastName  string    `gorm:"column:last_name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null"`
}

func (u *Users) TableName() string {
	return "users"
}
