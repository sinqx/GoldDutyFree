package service

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	db "sg/database"
	"sg/models"
	pb "sg/proto"
	"time"
)

var database = db.Connect()

func CreateUser(user *pb.User) (*pb.User, error) {
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "Nil pointer error")
	}

	if user.GetEmail() == "" || user.GetName() == "" ||
		user.GetPhoneNumber() == "" || user.GetAddress() == "" {
		return nil, fmt.Errorf("email name and phone number are empty")
	}

	newUser, err := fillDataBaseUserInfo(user)
	if err != nil {
		log.Println("Failed to create new User in Database: ", err)
	}

	database.NewRecord(newUser)
	database.Create(&newUser)
	user.Id = uint32(newUser.ID)
	user.CreatedAt = timestamppb.Now()
	user.UpdatedAt = timestamppb.Now()

	return user, nil
}

func GetUserById(id *pb.Id) (*pb.User, error) {
	var user models.User
	rowsAffected := database.Where("id = ?", id.Id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &pb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id")
	}

	return fillProtoUserInfo(&user)
}

func UpdateUserById(user *pb.User) (*pb.User, error) { // доделать проверку и изменения данных
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "Nil pointer error")
	}

	var updateUser models.User
	if err := database.First(&updateUser, user.GetId()).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "User with ID %v not found", user.GetId())
	}

	//if user.GetEmail() == "" || user.GetName() == "" ||
	//	user.GetPhoneNumber() == "" || user.GetAddress() == "" {
	//	return nil, status.Error(codes.InvalidArgument, "email name and phoneNumber are empty")
	//}

	if user.Name != "" && user.Name != updateUser.Name {
		updateUser.Name = user.Name
	}
	if user.Email != "" && user.Email != updateUser.Email {
		updateUser.Email = user.Email
	}
	if user.PhoneNumber != "" && user.PhoneNumber != updateUser.PhoneNumber {
		updateUser.PhoneNumber = user.PhoneNumber
	}
	if user.Address != "" && user.Address != updateUser.Address {
		updateUser.Address = user.Address
	}
	if user.Country != "" && user.Country != updateUser.Country {
		updateUser.Country = user.Country
	}
	if user.City != "" && user.City != updateUser.City {
		updateUser.City = user.City
	}
	if user.Zip != "" && user.Zip != updateUser.Zip {
		updateUser.Zip = user.Zip
	}

	updateUser.UpdatedAt = time.Now()
	if err := database.Save(&updateUser).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "Error updating user: %v", err)
	}

	return fillProtoUserInfo(&updateUser)
}

func GetAllUsers() (*pb.ListUser, error) {
	users := make([]*models.User, 0)
	err := database.Find(&users).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to get users from database")
	}

	pbUsers := make([]*pb.User, 0)
	for _, user := range users {
		pbUser, _ := fillProtoUserInfo(user)
		pbUsers = append(pbUsers, pbUser)
	}

	return &pb.ListUser{Users: pbUsers}, nil
}

func fillProtoUserInfo(userInfo *models.User) (*pb.User, error) {
	return &pb.User{
		Id:          uint32(userInfo.ID),
		Name:        userInfo.Name,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Address:     userInfo.Address,
		Country:     userInfo.Country,
		City:        userInfo.City,
		Zip:         userInfo.Zip,
		CreatedAt:   timestamppb.New(userInfo.CreatedAt),
		UpdatedAt:   timestamppb.New(userInfo.UpdatedAt),
	}, nil
}

func fillDataBaseUserInfo(userInfo *pb.User) (*models.User, error) {
	return &models.User{
		Name:        userInfo.Name,
		Email:       userInfo.Email,
		PhoneNumber: userInfo.PhoneNumber,
		Address:     userInfo.Address,
		Country:     userInfo.Country,
		City:        userInfo.City,
		Zip:         userInfo.Zip,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}, nil
}

//func CheckDuplicates(email, phone string) error {
//	var user models.User
//	if database.Where("email = ? OR phone_number = ?", email, phone).Find(&user).RecordNotFound() {
//		return nil
//	}
//	return fmt.Errorf("user with email %s or phone number %s already exists", email, phone)
//}
