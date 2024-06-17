package entity

import "laptopsShop/db"

type OrderStatus struct {
	Id     uint
	Status string
}

func (OrderStatus) TableName() string {
	return "order_status"
}

func MigrateOrderStatus() {
	if err := db.DB().AutoMigrate(&OrderStatus{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateOrderStatus)
}
