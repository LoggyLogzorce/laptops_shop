package entity

import "laptopsShop/db"

type Specifications struct {
	Id                         uint   `json:"id" gorm:"primaryKey"`
	Manufacturer               uint   `json:"manufacturer" gorm:"foreignKey:manufacturer"`
	Series                     string `json:"series"`
	Ram                        uint   `json:"ram" gorm:"foreignKey:ram"`
	MemoryFrequency            string `json:"memoryFrequency"`
	TypeRAM                    uint   `json:"typeRAM" gorm:"foreignKey:typeRAM"`
	ScreenSize                 string `json:"screenSize"`
	ScreenResolution           string `json:"screenResolution"`
	TypeScreen                 uint   `json:"typeScreen" gorm:"foreignKey:typeScreen"`
	TouchScreen                bool   `json:"touchScreen"`
	OperatingSystem            uint   `json:"operatingSystem" gorm:"foreignKey:operatingSystem"`
	Processor                  string `json:"processor"`
	NumberOfCores              uint   `json:"numberOfCores" gorm:"foreignKey:numberOfCores"`
	CacheMemory                uint   `json:"cacheMemory" gorm:"foreignKey:cacheMemory"`
	MaximumClocksSpeed         string `json:"maximumClocksSpeed"`
	ProcessorModel             string `json:"processorModel"`
	ManufacturerVideoProcessor uint   `json:"manufacturerVideoProcessor" gorm:"foreignKey:manufacturerVideoProcessor"`
	VideoCard                  string `json:"videoCard"`
	SSDCapacity                uint   `json:"SSDCapacity" gorm:"foreignKey:ssdCapacity"`
	Webcam                     string `json:"webcam"`
	WebcamQuality              string `json:"webcamQuality"`
	WiFiSupport                uint   `json:"wiFiSupport" gorm:"foreignKey:wifiSupport"`
	BluetoothVersion           uint   `json:"bluetoothVersion" gorm:"foreignKey:bluetoothVersion"`
	BatteryOperation           string `json:"batteryOperation"`
	Weight                     string `json:"weight"`
}

func (Specifications) TableName() string {
	return "specifications"
}

func MigrateSpecs() {
	if err := db.DB().AutoMigrate(&Specifications{}); err != nil {
		panic(err)
	}
}

func init() {
	db.Add(MigrateSpecs)
}
