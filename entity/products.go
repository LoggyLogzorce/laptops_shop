package entity

import "laptopsShop/db"

type Product struct {
	Id             uint   `json:"id" gorm:"primaryKey"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Price          uint   `json:"price"`
	Specifications uint   `json:"specifications" gorm:"foreignKey:specifications"`
	Image          string `json:"image"`
}

func (Product) TableName() string {
	return "products"
}

func MigrateProducts() {
	err := db.DB().AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateProducts)
}
