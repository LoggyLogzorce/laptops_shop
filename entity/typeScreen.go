package entity

import "laptopsShop/db"

type TypeScreen struct {
	Id         uint   `json:"id" gorm:"primaryKey"`
	TypeScreen string `json:"type-screen"`
}

func (TypeScreen) TableName() string {
	return "type_screen"
}

func MigrateTypeScreen() {
	if err := db.DB().AutoMigrate(&TypeScreen{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateTypeScreen)
}
