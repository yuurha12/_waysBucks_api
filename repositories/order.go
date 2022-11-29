package repositories

import (
	"waysbucks_BE/models"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindOrders() ([]models.Order, error)
	GetOrder(ID int) (models.Order, error)
	CreateOrder(order models.Order) (models.Order, error)
	// UpdateOrder(order models.Order) (models.Order, error)
	DeleteOrder(order models.Order) (models.Order, error)
	GetProductOrder(ID int) (models.Product, error)
	GetToppingOrder(ID []int) ([]models.Topping, error)
}

func RepositoryOrder(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindOrders() ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Buyer").Preload("Product").Preload("Topping").First(&orders).Error // Using Find method

	return orders, err
}

func (r *repository) GetOrder(ID int) (models.Order, error) {
	var order models.Order
	err := r.db.Preload("Buyer").Preload("Product").Preload("Topping").First(&order, ID).Error // Using First method

	return order, err
}

func (r *repository) CreateOrder(order models.Order) (models.Order, error) {
	err := r.db.Create(&order).Error // Using Create method

	return order, err
}

// func (r *repository) UpdateOrder(order models.Order) (models.Order, error) {
// 	err := r.db.Save(&order).Error // Using Save method

// 	return order, err
// }

func (r *repository) DeleteOrder(order models.Order) (models.Order, error) {
	err := r.db.Delete(&order).Error // Using Delete method

	return order, err
}

func (r *repository) GetProductOrder(ID int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, ID).Error
	return product, err
}

func (r *repository) GetToppingOrder(ID []int) ([]models.Topping, error) {
	var topping []models.Topping
	err := r.db.Find(&topping, ID).Error
	return topping, err
}
