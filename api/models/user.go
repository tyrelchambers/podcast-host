package models

import (
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

func CreateUser(user model.RegisterBody, db *gorm.DB) (u *model.User, e error) {

	id := cuid.New()

	var newUser model.User

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("ERROR: ", err)
		return u, err
	}

	db.Create(&model.User{
		UUID:     id,
		Email:    user.Email,
		Password: string(hashPassword),
	})

	fmt.Println("SUCCESS: new user created")
	return &newUser, nil
}

func GetUser(id string, db *gorm.DB) (user model.User, e error) {
	var u model.User

	db.Where("uuid = ?", id).First(&u)

	if db.Error != nil {
		return u, errors.New("Failed to get user. Doesn't exist.")
	}

	return u, nil
}

func FindUserByEmail(email string, db *gorm.DB) (user *model.User, e error) {
	var u model.User

	db.Where("email = ?", email).First(&u).Omit("password")

	if db.Error != nil {
		return nil, errors.New("Failed to get user. Doesn't exist.")
	}

	return &u, nil
}
