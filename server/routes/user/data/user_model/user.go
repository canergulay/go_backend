package user_model

type User struct {
	Id       int `gorm:"primaryKey"`
	Username string
	Mail     string
	Password string
	Updated  int64 `gorm:"autoUpdateTime:milli"`
	Created  int64 `gorm:"autoCreateTime"`
}

type userModel User
