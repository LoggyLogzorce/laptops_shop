package entity

import "laptopsShop/db"

type WiFiSupport struct {
	Id          uint   `json:"id" gorm:"primaryKey"`
	WiFiSupport string `json:"wifi_support"`
}

func (WiFiSupport) TableName() string {
	return "wifi_support"
}

func MigrateWifiSupport() {
	if err := db.DB().AutoMigrate(&WiFiSupport{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateWifiSupport)
}
