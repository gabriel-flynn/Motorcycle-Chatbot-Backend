package models

type Motorcycle struct {
	Id             uint    `gorm:"primaryKey;" json:"-"`
	Make           string  `json:"make"`
	Model          string  `json:"model"`
	YearStart      uint16  `json:"year_start"`
	YearEnd        uint16  `json:"year_end"`
	Price          float32 `json:"price"`
	Category       string  `json:"category"`
	EngineSize     string  `json:"engine_size"`
	EngineType     uint16  `json:"engine_type"`
	InsuranceGroup uint8   `json:"insurance_group"`
	MPG            uint8   `json:"mpg"`
	TankRange      uint16  `json:"tank_range"`
	Power          uint16  `json:"power"`
	SeatHeight     uint8   `json:"seat_height"`
	Weight         uint16  `json:"weight"`
	Review         *Review `json:"review"`
	ReviewId       uint    `json:"-"`
}
