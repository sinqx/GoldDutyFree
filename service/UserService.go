package service

import (
	_ "github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "sg/database"
	"sg/models"

	pb "sg/proto"
)

var database = db.Connect()

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func CreateUser(user *pb.User) (*pb.User, error) {
	newUser := models.User{Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}

	database.NewRecord(newUser)
	database.Create(&newUser)

	return &pb.User{Id: user.Id, Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}

func findUserById(id uint32) (*pb.User, error) { //????????
	var user models.User
	rowsAffected := database.Where("id = ?", id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &pb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id!")
	}

	return &pb.User{Id: uint32(user.ID), Name: user.Name, Email: user.Email, PhoneNumber: user.PhoneNumber}, nil
}
