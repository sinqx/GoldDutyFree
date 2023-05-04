package service

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"sg/models"
	pb "sg/proto"
	"time"
)

func CreateOrder(OrderModel *pb.OrderModel) (*pb.Order, error) {

	user, err := GetUserById(&pb.Id{Id: OrderModel.UserId})
	if err != nil {
		log.Fatalf("Person with this id not found: %v", err)
	}

	price, err := GetGoldPrice(OrderModel.Gold)
	if err != nil {
		log.Fatalf("Gold Service error: %v", err)
	}

	var dbUser models.User
	rowsAffected := database.Where("id = ?", &OrderModel.UserId).Find(&user).RowsAffected
	if rowsAffected == 0 {
		log.Fatalf("Cannot find a User with this id: %v", err)
	}

	var dbGold = models.GoldModel{
		Weight:    OrderModel.Gold.Weight,
		Quantity:  uint32(OrderModel.Gold.Quantity),
		PriceTime: time.Time{},
	}

	dbOrder := models.Order{
		Price:       uint(price.Price),
		Gold:        dbGold,
		User:        dbUser,
		OrderStatus: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	database.NewRecord(dbOrder)
	database.Create(&dbOrder)

	return &pb.Order{
		Id:          OrderModel.Id,
		Price:       price,
		User:        user,
		Gold:        OrderModel.Gold,
		OrderStatus: 0,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	}, nil
}
