package _interface

import (
	"context"
	pb "sg/proto"
	"sg/service"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
}
type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *UserServer) CreateUser(_ context.Context, User *pb.User) (*pb.User, error) {
	return service.CreateUser(User)
}

func (s *UserServer) UpdateUserById(_ context.Context, User *pb.User) (*pb.User, error) {
	return service.UpdateUserById(User)
}

func (s *UserServer) FindUserById(_ context.Context, Id *pb.Id) (*pb.User, error) {
	return service.GetUserById(Id)
}

func (s *UserServer) GetAllUsers(_ context.Context, _ *pb.User) (*pb.ListUser, error) {
	return service.GetAllUsers()
}

func (s *OrderServer) CreateOrder(_ context.Context, OrderModel *pb.OrderModel) (*pb.Order, error) {
	return service.CreateOrder(OrderModel)
}

func (s *OrderServer) GetAllOrders(_ context.Context, _ *pb.Order) (*pb.ListOrder, error) {
	return service.GetAllOrders()
}

func (s *OrderServer) GetGoldPrice(_ context.Context, Gold *pb.Gold) (*pb.Price, error) {
	return service.GetGoldPrice(Gold)
}
