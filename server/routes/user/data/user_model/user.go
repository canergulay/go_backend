package user_model

type User struct {
	ID       int `gorm:"primaryKey;AUTO_INCREMENT"`
	Username string
	Mail     string
	Password string
	Updated  int64 `gorm:"autoUpdateTime:milli"`
	Created  int64 `gorm:"autoCreateTime"`
}
