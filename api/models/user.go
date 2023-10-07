package models

import (
	"api/helpers"
	"api/model"
	"errors"
	"fmt"

	"github.com/lucsky/cuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func CheckUserExists(email string, db *gorm.DB) (userExists bool) {
	if email == "" {
		panic(errors.New("Email cannot be empty."))
	}

	db.Where("email = ?", email).First(&model.User{})

	if db.Error != nil {
		return false
	}

	if db.RowsAffected == 0 {
		return false
	}

	return true
}

func CreateUser(user model.RegisterBody, db *gorm.DB) (*model.User, error) {

	id := cuid.New()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	newUser := model.User{
		UUID:     id,
		Email:    user.Email,
		Password: string(hashPassword),
	}

	db.Create(&newUser)

	if db.Error != nil {
		fmt.Println("ERROR: ", err)
		return nil, err
	}

	fmt.Printf("SUCCESS: new user created: %s\n", newUser.UUID)
	return &newUser, nil
}

func GetUser(id string, db *gorm.DB) (*model.UserDTO, error) {
	var u model.User
	var uDto model.UserDTO

	db.Table("users").Where("uuid = ?", id).First(&u)

	if db.Error != nil {
		return nil, errors.New("Failed to get user. Doesn't exist.")
	}

	helpers.ConvertToDto(u, &uDto)

	return &uDto, nil
}

func FindUserByEmail(email string, db *gorm.DB) (*model.User, error) {
	var u model.User

	db.Table("users").Where("email = ?", email).First(&u).Omit("password")

	if db.Error != nil {
		return nil, errors.New("Failed to get user. Doesn't exist.")
	}

	return &u, nil
}
