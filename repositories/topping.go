package repositories

import (
	"waysbucks_BE/models"

	"gorm.io/gorm"
)

type ToppingRepository interface {
	FindToppings() ([]models.Topping, error)
	GetTopping(ID int) (models.Topping, error)
	CreateTopping(topping models.Topping) (models.Topping, error)
	DeleteTopping(topping models.Topping) (models.Topping, error)
	UpdateTopping(topping models.Topping) (models.Topping, error)
}

func RepositoryTopping(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindToppings() ([]models.Topping, error) {
	var Toppings []models.Topping
	err := r.db.Preload("User").Find(&Toppings).Error

	return Toppings, err
}

func (r *repository) GetTopping(ID int) (models.Topping, error) {
	var topping models.Topping
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.Preload("User").First(&topping, ID).Error

	return topping, err
}

func (r *repository) CreateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Create(&topping).Error

	return topping, err
}

func (r *repository) DeleteTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Delete(&topping).Error

	return topping, err
}

func (r *repository) UpdateTopping(topping models.Topping) (models.Topping, error) {
	err := r.db.Save(&topping).Error // Using Save method

	return topping, err
}
