package entity

import (
	"github.com/google/uuid"
	"laptopsShop/db"
)

type Order struct {
	Id         uint
	User       uuid.UUID
	Status     uint
	TotalPrice float32
}

func (Order) TableName() string {
	return "orders"
}

func MigrateOrders() {
	if err := db.DB().AutoMigrate(&Order{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateOrders)
}
