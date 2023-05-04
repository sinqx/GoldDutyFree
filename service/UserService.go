package service

import (
	_ "github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	db "sg/database"
	"sg/models"
	pb "sg/proto"
	"time"
)

var database = db.Connect()

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func CreateUser(user *pb.User) (*pb.User, error) {
	newUser := models.User{
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Country:     user.Country,
		City:        user.City,
		Zip:         user.Zip,
		CVV:         user.Cvv,
		CreatedAt:   time.Now(),
	}

	database.NewRecord(newUser)
	database.Create(&newUser)

	return user, nil
}

func GetUserById(id *pb.Id) (*pb.User, error) { //????????
	var user models.User
	rowsAffected := database.Where("id = ?", id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &pb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id")
	}

	return &pb.User{
		Id:          uint32(user.ID),
		Name:        user.Name,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Country:     user.Country,
		City:        user.City,
		Zip:         user.Zip,
		Cvv:         user.CVV,
	}, nil
}

func UpdateUserById(user *pb.User) (*pb.User, error) {

	id := user.GetId()
	name := user.GetName()
	email := user.GetEmail()
	phoneNumber := user.GetPhoneNumber()

	var updateUser models.User
	database.Where("id =?", id).Find(&updateUser)

	updateUser.Name = name
	updateUser.Email = email
	updateUser.PhoneNumber = phoneNumber

	database.Save(&updateUser)

	return &pb.User{Id: uint32(updateUser.ID), Name: updateUser.Name, Email: updateUser.Email, PhoneNumber: updateUser.PhoneNumber}, nil
}
