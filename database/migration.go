package database

import (
	"fmt"
	"ways-bucks-api/models"
	"ways-bucks-api/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Topping{}, &models.Profile{}, &models.Cart{}, &models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
