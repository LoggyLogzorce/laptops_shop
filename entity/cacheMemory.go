package entity

import "laptopsShop/db"

type CacheMemory struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	CacheMemory string `json:"cache_memory"`
}

func (CacheMemory) TableName() string {
	return "cache_memory"
}

func MigrateCacheMemory() {
	if err := db.DB().AutoMigrate(&CacheMemory{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateCacheMemory)
}
