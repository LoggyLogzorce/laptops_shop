package entity

import "laptopsShop/db"

type TypeRAM struct {
	Id      uint   `json:"id" gorm:"primaryKey"`
	TypeRAM string `json:"type-ram"`
}

func (TypeRAM) TableName() string {
	return "type_ram"
}

func MigrateTypeRAM() {
	if err := db.DB().AutoMigrate(&TypeRAM{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateTypeRAM)
}
