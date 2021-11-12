package user_model

type User struct {
	ID                 int    `gorm:"primaryKey;AUTO_INCREMENT"`
	Username           string `gorm:"not null"`
	Mail               string `gorm:"not null"`
	Password           string `gorm:"not null"`
	Updated            int64  `gorm:"autoUpdateTime:milli"`
	Created            int64  `gorm:"autoCreateTime"`
	RegisterMethod     string
	IsEmailValidated   bool `gorm:"default:false"`
	IsOnboardCompleted bool `gorm:"default:false"`
}
