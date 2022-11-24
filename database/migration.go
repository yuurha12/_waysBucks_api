package database

import (
	"fmt"
	"waysbucks_BE/models"
	"waysbucks_BE/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{},
		// &models.Profile{},
		&models.Product{},
		// &models.Topping{},
		&models.Transaction{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}