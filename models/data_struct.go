package models

import (
	pb "sg/proto"
	"time"
)

type User struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Address     string `json:"address" form:"Address"`
	Country     string `json:"country" form:"Country"`
	City        string `json:"city" form:"City"`
	Zip         string `json:"zip" form:"Zip"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Gold struct {
	ID        uint          `gorm:"primary_key"`
	Weight    pb.GoldWeight `json:"weight"`
	Quantity  float32       `json:"quantity"`
	PriceTime time.Time
}

type Order struct {
	ID          uint               `json:"ID,omitempty"`
	Price       uint32             `json:"price"`
	UserID      uint               `json:"userID"`
	GoldID      uint               `json:"goldID"`
	User        User               `json:"user" gorm:"foreignKey:UserID"`
	Gold        Gold               `json:"gold" gorm:"foreignKey:GoldID"`
	OrderStatus pb.OrderStatusEnum `json:"orderStatus,omitempty"`
	CreatedAt   time.Time          `json:"createdAt"`
	UpdatedAt   time.Time          `json:"updatedAt"`
}

//
//type GoldWeight struct {
//	ounce string
//	gram  string
//	kg    string
//}
//
//type OrderStatusEnum struct {
//	inProgress string
//	Confirmed  string
//	Canceled   string
//	Delivered  string
//	Delivering string
//}
