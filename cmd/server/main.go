package main

import (
	"money-manager/internal/db"
	"money-manager/internal/models"
	"money-manager/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	// os.Setenv("DB_URL", "postgres://postgres:mysecretpassword@localhost:5433/money_manager")

	db.Connect()

	db.DB.AutoMigrate(
		&models.Account{},
		&models.Transaction{},
	)

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}