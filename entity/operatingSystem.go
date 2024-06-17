package entity

import "laptopsShop/db"

type OperatingSystem struct {
	Id              uint   `json:"id" gorm:"primaryKey"`
	OperatingSystem string `json:"operating_system"`
}

func (OperatingSystem) TableName() string {
	return "operating_system"
}

func MigrateOperatingSystem() {
	if err := db.DB().AutoMigrate(&OperatingSystem{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateOperatingSystem)
}
