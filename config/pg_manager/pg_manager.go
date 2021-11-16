package pg_manager

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func InitPostgreSQL() *gorm.DB {
	dsn := connectionStringGetter()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("wE GOT AA!")
	}
	connection = db
	return connection

}

func GetPostgresConnection() *gorm.DB {
	return connection
}

func connectionStringGetter() string {
	server := os.Getenv("SERVER")
	user := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	return "host=" + server + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"

}
