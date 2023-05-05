package service

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "sg/proto"
	"testing"
)

func TestCreateUser(t *testing.T) {

	testUsers := make([]*pb.User, 0)

	testUsers = append(testUsers,
		&pb.User{
			Id:          6,
			Name:        "name",
			Email:       "test@test.com",
			PhoneNumber: "123456789",
			Address:     "Test Address",
			Country:     "KG",
			City:        "BSHK",
			Zip:         "12345",
			Cvv:         "223",
			CreatedAt:   timestamppb.Now(),
			UpdatedAt:   timestamppb.Now(),
		},
		&pb.User{
			Id:          0,
			Name:        "",
			Email:       "",
			PhoneNumber: "",
			Address:     "",
			Country:     "",
			City:        "",
			Zip:         "",
			Cvv:         "",
			CreatedAt:   nil,
			UpdatedAt:   nil,
		},
		&pb.User{

			Name:        "fffff",
			Email:       "fffff",
			PhoneNumber: "ffff",
			Address:     "",
			Country:     "",

			Zip: "",
			Cvv: "",

			UpdatedAt: nil,
		})

	for _, user := range testUsers {
		createdUser, err := CreateUser(user)

		// Проверяем на nil
		assert.NotNil(t, createdUser)

		assert.NotEmpty(t, createdUser.Id)
		assert.NotEmpty(t, createdUser.Email)

		// Проверяем, что ошибка при создании юзера не возникла
		assert.NoError(t, err)
	}

}
