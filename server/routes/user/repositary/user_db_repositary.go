package repositary

import (
	"backend/server/routes/user/model"
	"backend/server/routes/user/utils"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserDbRepositary struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserDbRepositary {
	return &UserDbRepositary{
		db: db,
	}
}

func (r *UserDbRepositary) CheckIfUserExists(mail string) bool {
	var user model.User
	result := r.db.Where("mail = ?", mail).First(&user)
	return result.Error == nil
}

func (r *UserDbRepositary) LoginUser(mail string, password string) model.LoginResponseModel {
	var user model.User

	result := r.db.Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return model.LoginResponseModel{Status: NotAnAccount}
	}
	isPaswordsMatched := utils.ComparePasswords(user.Password, password)
	if isPaswordsMatched {
		return model.LoginResponseModel{Status: Succes, User: user}

	}
	return model.LoginResponseModel{Status: WrongPass}
}

func (r *UserDbRepositary) RegisterUser(mail string, username string, password string) (model.RegisterResponseModel, error) {
	hashedPasword, err := utils.HashMyPassword(password)
	if err != nil {
		return model.RegisterResponseModel{IsRegistrationCompleted: false}, err
	}
	loginResponse := r.LoginUser(mail, password)
	if loginResponse.Status != NotAnAccount {
		// IF THERE EXIST AN ACCOUNT WITH THAT EMAIL
		return model.RegisterResponseModel{IsRegistrationCompleted: false}, errors.New(alreadyRegistered)
	}
	userToRegister := model.User{Username: username, Mail: mail, Password: hashedPasword, RegisterMethod: "email"}

	result := r.db.Create(&userToRegister)
	if result.Error != nil {
		fmt.Println("we got a problem")
		fmt.Println(result.Error)
		return model.RegisterResponseModel{IsRegistrationCompleted: false}, result.Error

	}
	return model.RegisterResponseModel{IsRegistrationCompleted: true, User: userToRegister}, nil

}

const (
	Succes int = iota
	WrongPass
	NotAnAccount
)

const alreadyRegistered = "there is already an account with that email"
