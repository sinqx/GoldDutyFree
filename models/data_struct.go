package models

import (
	"github.com/jinzhu/gorm"
	pb "sg/proto"
	"time"
)

type Order struct {
	ID           uint               `gorm:"primary_key"`
	Price        uint               `json:"price"`
	GoldWeight   uint               `json:"GoldWeight"`
	GoldQuantity uint               `json:"GoldQuantity"`
	User         User               `json:"user"`
	OrderStatus  pb.OrderStatusEnum `json:"orderStatus"`
	CreatedAt    time.Time
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
	ID           uint `gorm:"primary_key"`
	UserId       uint `json:"userId" form:"userId"`
	GoldWeight   uint `json:"goldWeight" form:"goldWeight"`
	GoldQuantity uint `json:"goldQuantity" form:"goldQuantity"`
}
