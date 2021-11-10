package user_db

import (
	"backend/config/pg_manager"
	pw_hasher "backend/server/routes/user/data/pw_hasher"
	"backend/server/routes/user/data/user_model"
	"fmt"
)

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

func LoginUser(mail string, password string) int {
	var user user_model.User
	db := pg_manager.GetPostgresConnection()
	result := db.Where("mail = ?", mail).First(&user)
	if result.Error != nil {
		return NotAnAccount
	}
	isPaswordsMatched := pw_hasher.ComparePasswords(user.Password, password)
	if isPaswordsMatched {
		return Succes
	}
	return WrongPass
}

func RegisterUser(mail string, username string, password string) (bool, error) {
	hashedPasword, err := pw_hasher.HashMyPassword(password)
	if err != nil {
		return false, err
	}
	userToRegister := user_model.User{Username: username, Mail: mail, Password: hashedPasword}
	db := pg_manager.GetPostgresConnection()
	result := db.Create(&userToRegister)
	if result.Error != nil {
		fmt.Println("we got a problem")
		fmt.Println(result.Error)
		return false, result.Error

	}
	return true, nil

}

const (
	Succes int = iota
	WrongPass
	NotAnAccount
)
