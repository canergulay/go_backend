package user_db

import (
	"backend/config/pg_manager"
	"backend/server/routes/user/data/user_model"
	"fmt"
)

func AutoMigrate() {
	db := pg_manager.GetPostgresConnection()
	db.AutoMigrate(&user_model.User{})
}

func CheckIfUserExists(mail string) {
	db := pg_manager.GetPostgresConnection()
	user := db.First(&user_model.User{Mail: mail})
	if user != nil {
		fmt.Println("yeah we found it!")
	} else {
		fmt.Println("no such user!")
	}

}
