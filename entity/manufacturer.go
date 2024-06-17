package entity

import "laptopsShop/db"

type Manufacturer struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Manufacturer string `json:"manufacturer"`
}

func (Manufacturer) TableName() string {
	return "manufacturer"
}

func MigrateManufacturer() {
	err := db.DB().AutoMigrate(&Manufacturer{})
	if err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateManufacturer)
}
