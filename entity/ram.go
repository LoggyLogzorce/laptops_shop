package entity

import "laptopsShop/db"

type Ram struct {
	Id  uint   `json:"id" gorm:"primaryKey"`
	Ram string `json:"ram"`
}

func (Ram) TableName() string {
	return "ram"
}

func MigrateRam() {
	if err := db.DB().AutoMigrate(&Ram{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateRam)
}
