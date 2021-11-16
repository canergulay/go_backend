package user

import (
	"backend/config/pg_manager"
	"backend/pkg/routes/user/data/pw_hasher"
	"backend/pkg/routes/user/model"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type UserDbRepositary struct {
	db *gorm.DB
}

func NewDBRepository(db *gorm.DB) *UserDbRepositary {
	return &UserDbRepositary{
		db: db,
	}
}

func AutoMigrate() {
	db := pg_manager.GetPostgresConnection()
	db.AutoMigrate(&model.User{})
}

func CheckIfUserExists(mail string) bool {
	var user model.User
	db := pg_manager.GetPostgresConnection()
	result := db.Where("mail = ?", mail).First(&user)
	return result.Error == nil
}

func LoginUser(mail string, password string) model.LoginResponseModel {
	var user model.User
	db := pg_manager.GetPostgresConnection()
	result := db.Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return model.LoginResponseModel{Status: NotAnAccount}
	}
	isPaswordsMatched := pw_hasher.ComparePasswords(user.Password, password)
	if isPaswordsMatched {
		return model.LoginResponseModel{Status: Succes, User: user}

	}
	return model.LoginResponseModel{Status: WrongPass}
}

func RegisterUser(mail string, username string, password string) (model.RegisterResponseModel, error) {
	hashedPasword, err := pw_hasher.HashMyPassword(password)
	if err != nil {
		return model.RegisterResponseModel{IsRegistrationCompleted: false}, err
	}
	loginResponse := LoginUser(mail, password)
	if loginResponse.Status != NotAnAccount {
		// IF THERE EXIST AN ACCOUNT WITH THAT EMAIL
		return model.RegisterResponseModel{IsRegistrationCompleted: false}, errors.New(alreadyRegistered)
	}
	userToRegister := model.User{Username: username, Mail: mail, Password: hashedPasword, RegisterMethod: "email"}
	db := pg_manager.GetPostgresConnection()
	result := db.Create(&userToRegister)
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
