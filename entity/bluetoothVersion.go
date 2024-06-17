package entity

import "laptopsShop/db"

type BluetoothVersion struct {
	Id               uint   `json:"id" gorm:"primaryKey"`
	BluetoothVersion string `json:"bluetooth_version"`
}

func (BluetoothVersion) TableName() string {
	return "bluetooth_version"
}

func MigrateBluetoothVersion() {
	if err := db.DB().AutoMigrate(&BluetoothVersion{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateBluetoothVersion)
}
