package database

import (
	"fmt"
	"waysbuck/models"
	"waysbuck/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.User{}, 
		&models.Product{}, 
		&models.Topping{}, 
		&models.Order{}, 
		&models.Transaction{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}