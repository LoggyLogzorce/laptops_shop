package entity

import "laptopsShop/db"

type ProductsInOrders struct {
	Id         uint
	Order      uint
	Product    uint
	Quantity   uint
	TotalPrice float32
}

func (ProductsInOrders) TableName() string {
	return "products_in_orders"
}

func MigrateProductsInOrders() {
	if err := db.DB().AutoMigrate(&ProductsInOrders{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateProductsInOrders)
}
