package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID       uint   `gorm:"not null" json:"userId"`
	ProductName  string `gorm:"not null" json:"productName"`
	QuantityType string `gorm:"not null" json:"quantityType"`
	Quantity     int    `gorm:"not null" json:"quantity"`
}

func AddOrder(user *User, order *Order) error {
	db.Model(&user).
		Preload("Orders", "email = ?", user.Email).
		Association("Orders").
		Append(order)
	return db.Save(&order).Error
}

func GetOrders(user *User) (*[]Order, error) {
	var orders *[]Order
	err := db.Model(&user).
		Association("Orders").
		Find(&orders)
	return orders, err
}
