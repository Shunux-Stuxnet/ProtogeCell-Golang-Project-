package models

type User struct {
	ID    string `gorm:"primary" json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Device struct {
	IMEI  int `gorm:"unique, primaryKey" json:"imei"`
	Count int `json:"count" gorm:"default:0"`
}

type IMEI struct {
	IMEI     int64  `gorm:"primaryKey" json:"imei" form:"imei"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	FIR      string `json:"fir" form:"fir"`
	Mobile   string `json:"phone" form:"mobile"`
	Location string `json:"location" form:"location"`
	Info     string `json:"message" form:"info"`
}
