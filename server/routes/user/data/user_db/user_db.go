package user_db

import (
	"backend/config/pg_manager"
	"backend/server/routes/user/data/user_model"
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
