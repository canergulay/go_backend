package model

import "gorm.io/gorm"

type User struct {
	ID                 int    `json:"id" gorm:"primaryKey;AUTO_INCREMENT"`
	Username           string `json:"username" gorm:"not null"`
	Mail               string `json:"mail" gorm:"not null"`
	Picture            string `json:"picture" gorm:"not null"`
	Password           string `json:"password" gorm:"not null"`
	Updated            int64  `json:"updatedAt" gorm:"autoUpdateTime:milli"`
	Created            int64  `json:"createdAt" gorm:"autoCreateTime"`
	RegisterMethod     string `json:"registermethod" gorm:"not null"`
	IsEmailValidated   bool   `json:"isEmailValidated" gorm:"default:false"`
	IsOnboardCompleted bool   `json:"isOnboardCompleted" gorm:"default:false"`
}

func AutoMigrateUserModel(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
