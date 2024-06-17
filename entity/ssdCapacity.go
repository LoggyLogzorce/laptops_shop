package entity

import "laptopsShop/db"

type SSDCapacity struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	SSDCapacity string `json:"ssd_capacity"`
}

func (SSDCapacity) TableName() string {
	return "ssd_capacity"
}

func MigrateSSDCapacity() {
	if err := db.DB().AutoMigrate(&SSDCapacity{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateSSDCapacity)
}
