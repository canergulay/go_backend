package model

import "gorm.io/gorm"

type User struct {
	ID                 int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username           string `gorm:"not null"`
	Mail               string `gorm:"not null"`
	Password           string `gorm:"not null"`
	Updated            int64  `gorm:"autoUpdateTime:milli"`
	Created            int64  `gorm:"autoCreateTime"`
	RegisterMethod     string `gorm:"not null"`
	IsEmailValidated   bool   `gorm:"default:false"`
	IsOnboardCompleted bool   `gorm:"default:false"`
}

func AutoMigrateUserModel(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
