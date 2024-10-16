package main

import (
	"fmt"

	"github.com/thanpawatpiti/exec-go-loan/config"
	"github.com/thanpawatpiti/exec-go-loan/config/database"
	"github.com/thanpawatpiti/exec-go-loan/pkg/cmd"
	"github.com/thanpawatpiti/exec-go-loan/pkg/models"
	"github.com/thanpawatpiti/exec-go-loan/routes"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	fmt.Println("init")
	config.Init()
	db = database.NewConnector().Connect()
	if db == nil {
		panic("Connot connect to db")
	}

	fmt.Println("Connect DB success")
}

func main() {
	fmt.Println("main")

	// Init Model
	model := models.NewInit(db)

	// Init Command
	cmd := cmd.NewInitCmd(model)

	// Load routes
	r := routes.NewRoute(cmd, model)
	r.Load()

	// Start the server
	r.Start()
}
