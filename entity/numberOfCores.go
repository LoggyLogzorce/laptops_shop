package entity

import "laptopsShop/db"

type NumberOfCores struct {
	Id            uint   `json:"id" gorm:"primaryKey"`
	NumberOfCores string `json:"number_of_cores"`
}

func (NumberOfCores) TableName() string {
	return "number_of_cores"
}

func MigrateNumberOfCores() {
	if err := db.DB().AutoMigrate(&NumberOfCores{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateNumberOfCores)
}
