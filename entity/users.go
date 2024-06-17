package entity

import (
	"github.com/google/uuid"
	"laptopsShop/db"
)

type Users struct {
	Id       uuid.UUID `json:"id" gorm:"primaryKey"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}

func (Users) TableName() string {
	return "users"
}

func MigrateUsers() {
	if err := db.DB().AutoMigrate(&Users{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateUsers)
}
