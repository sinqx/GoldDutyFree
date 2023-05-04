package service

import (
	pb "sg/proto"
	//"UserService"
)

func CreateOrder(OrderModel *pb.OrderModel) (pb.Order, error) {

	user, _ := GetUserById(&OrderModel.UserId)
	return
}
