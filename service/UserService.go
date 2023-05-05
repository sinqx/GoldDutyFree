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
		return nil, fmt.Errorf("email name and phoneNumber are empty")
	}

	newUser, err := fillDataBaseUserInfo(user)
	if err != nil {
		log.Println("Failed to create new User in Database: ", err)
	}

	database.NewRecord(newUser)
	database.Create(&newUser)
	user.Id = uint32(newUser.ID)

	return user, nil
}

func GetUserById(id *pb.Id) (*pb.User, error) { //????????
	var user models.User
	rowsAffected := database.Where("id = ?", id.Id).Find(&user).RowsAffected

	if rowsAffected == 0 {
		return &pb.User{}, status.Error(codes.NotFound, "Cannot find a User with this id")
	}

	return fillProtoUserInfo(&user)
}

func UpdateUserById(user *pb.User) (*pb.User, error) {
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "Nil pointer error")
	}
	if user.GetEmail() == "" || user.GetName() == "" ||
		user.GetPhoneNumber() == "" || user.GetAddress() == "" {
		return nil, status.Error(codes.InvalidArgument, "email name and phoneNumber are empty")
	}

	var updateUser models.User
	database.Where("id =?", user.GetId()).Find(&updateUser)

	updateUser.Email = user.GetEmail()
	updateUser.PhoneNumber = user.GetPhoneNumber()
	updateUser.Address = user.GetAddress()
	updateUser.Zip = user.GetZip()
	updateUser.UpdatedAt = time.Now()

	database.Save(&updateUser)

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
		Cvv:         userInfo.CVV,
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
		CVV:         userInfo.Cvv,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

//func CheckDuplicates(email, phone string) error {
//	var user models.User
//	if database.Where("email = ? OR phone_number = ?", email, phone).Find(&user).RecordNotFound() {
//		return nil
//	}
//	return fmt.Errorf("user with email %s or phone number %s already exists", email, phone)
//}
