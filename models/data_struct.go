package models

import (
	"github.com/jinzhu/gorm"
	pb "sg/proto"
	"time"
)

type Order struct {
	ID          uint               `gorm:"primary_key"`
	Price       uint               `json:"price"`
	Gold        GoldModel          `json:"gold"`
	User        User               `json:"user"`
	OrderStatus pb.OrderStatusEnum `json:"orderStatus"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Address     string `json:"address" form:"Address"`
	Country     string `json:"country" form:"Country"`
	City        string `json:"city" form:"City"`
	Zip         string `json:"zip" form:"Zip"`
	CVV         string `json:"CVV" form:"CVV"`
	CreatedAt   time.Time
}

type OrderModel struct {
	gorm.Model
	UserId       uint          `json:"userId" form:"userId"`
	GoldWeight   pb.GoldWeight `json:"goldWeight" form:"goldWeight"`
	GoldQuantity uint          `json:"goldQuantity" form:"goldQuantity"`
}

type GoldModel struct {
	gorm.Model
	Weight    pb.GoldWeight `json:"weight"`
	Quantity  uint32        `json:"quantity"`
	PriceTime time.Time
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
