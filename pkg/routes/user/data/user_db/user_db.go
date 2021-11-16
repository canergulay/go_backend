package user_db

import (
	"backend/config/pg_manager"
	pw_hasher "backend/pkg/routes/user/data/pw_hasher"
	"backend/pkg/routes/user/data/user_model"
	"errors"
	"fmt"
)

type LoginResponseModel struct {
	Status int
	User   user_model.User
}

type RegisterResponseModel struct {
	IsRegistrationCompleted bool
	User                    user_model.User
}

func AutoMigrate() {
	db := pg_manager.GetPostgresConnection()
	db.AutoMigrate(&user_model.User{})
}

func CheckIfUserExists(mail string) bool {
	var user user_model.User
	db := pg_manager.GetPostgresConnection()
	result := db.Where("mail = ?", mail).First(&user)
	return result.Error == nil
}

func LoginUser(mail string, password string) LoginResponseModel {
	var user user_model.User
	db := pg_manager.GetPostgresConnection()
	result := db.Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return LoginResponseModel{Status: NotAnAccount}
	}
	isPaswordsMatched := pw_hasher.ComparePasswords(user.Password, password)
	if isPaswordsMatched {
		return LoginResponseModel{Status: Succes, User: user}

	}
	return LoginResponseModel{Status: WrongPass}
}

func RegisterUser(mail string, username string, password string) (RegisterResponseModel, error) {
	hashedPasword, err := pw_hasher.HashMyPassword(password)
	if err != nil {
		return RegisterResponseModel{IsRegistrationCompleted: false}, err
	}
	loginResponse := LoginUser(mail, password)
	if loginResponse.Status != NotAnAccount {
		// IF THERE EXIST AN ACCOUNT WITH THAT EMAIL
		return RegisterResponseModel{IsRegistrationCompleted: false}, errors.New(alreadyRegistered)
	}
	userToRegister := user_model.User{Username: username, Mail: mail, Password: hashedPasword, RegisterMethod: "email"}
	db := pg_manager.GetPostgresConnection()
	result := db.Create(&userToRegister)
	if result.Error != nil {
		fmt.Println("we got a problem")
		fmt.Println(result.Error)
		return RegisterResponseModel{IsRegistrationCompleted: false}, result.Error

	}
	return RegisterResponseModel{IsRegistrationCompleted: true, User: userToRegister}, nil

}

const (
	Succes int = iota
	WrongPass
	NotAnAccount
)

const alreadyRegistered = "there is already an account with that email"
