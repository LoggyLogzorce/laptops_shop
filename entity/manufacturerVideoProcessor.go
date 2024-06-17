package entity

import "laptopsShop/db"

type ManufacturerVideoProcessor struct {
	Id                         uint   `json:"id" gorm:"primaryKey"`
	ManufacturerVideoProcessor string `json:"manufacturer_video_processor"`
}

func (ManufacturerVideoProcessor) TableName() string {
	return "manufacturer_video_processor"
}

func MigrateManufacturerVideoProcessor() {
	if err := db.DB().AutoMigrate(&ManufacturerVideoProcessor{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateManufacturerVideoProcessor)
}
