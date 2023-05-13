package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"sg/models"
	pb "sg/proto"
	"time"
)

func CreateOrder(OrderModel *pb.OrderModel) (*pb.Order, error) {
	user, err := GetUserById(&pb.Id{Id: OrderModel.UserId})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Cannot find a User with this id: %v", err)
	}

	price, err := GetGoldPrice(OrderModel.Gold)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Gold Service error: %v", err)
	}

	var dbUser models.User
	database.Where("id = ?", &OrderModel.UserId).Find(&dbUser)
	if dbUser.ID == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find a User with this id")
	}

	dbOrder := models.Order{
		Price:  uint32(price.Price),
		UserID: dbUser.ID,
		GoldID: 0,
		User:   dbUser,
		Gold: models.Gold{
			Weight:    OrderModel.Gold.Weight,
			Quantity:  OrderModel.Gold.Quantity,
			PriceTime: time.Now(),
		},
		OrderStatus: 0,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	database.NewRecord(dbOrder)
	database.Create(&dbOrder)

	OrderModel.Gold.PriceTime = timestamppb.Now()
	return &pb.Order{
		Price:       uint32(price.Price),
		User:        user,
		Gold:        OrderModel.Gold,
		OrderStatus: 0,
		CreatedAt:   timestamppb.Now(),
		UpdatedAt:   timestamppb.Now(),
	}, nil
}

func GetAllOrders() (*pb.ListOrder, error) {
	orders := make([]*models.Order, 0)
	err := database.Find(&orders).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get users from database")
	}

	pbOrders := make([]*pb.Order, 0)
	for _, dbOrder := range orders {
		user, _ := fillProtoUserInfo(&dbOrder.User)

		gold := &pb.Gold{
			Weight:    dbOrder.Gold.Weight,
			Quantity:  float32(dbOrder.Gold.Quantity),
			PriceTime: timestamppb.New(dbOrder.Gold.PriceTime),
		}

		order := &pb.Order{
			Id:          uint32(dbOrder.ID),
			Price:       dbOrder.Price,
			User:        user,
			Gold:        gold,
			OrderStatus: dbOrder.OrderStatus,
			CreatedAt:   timestamppb.New(dbOrder.CreatedAt),
			UpdatedAt:   timestamppb.New(dbOrder.UpdatedAt),
		}
		pbOrders = append(pbOrders, order)
	}

	return &pb.ListOrder{Orders: pbOrders}, nil
}
