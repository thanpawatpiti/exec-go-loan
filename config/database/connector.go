package database

import (
	"fmt"
	"time"

	config "github.com/thanpawatpiti/exec-go-loan/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnector() *Connector {
	return &Connector{}
}

type Connector struct {
}

// ConnectDBWithCaCert connects to the database using a CA certificate.
func (c *Connector) Connect() (db *gorm.DB) {
	// Create the DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName,
	)

	// Connect to the database using gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NowFunc: func() time.Time {
			loc, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(loc)
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	return db
}
