package user

import (
	"backend/global/authentication"
	"backend/server/routes/user/api"
	"backend/server/routes/user/model"
	"backend/server/routes/user/repositary"
	"backend/server/routes/user/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(r *gin.Engine, db *gorm.DB) {
	model.AutoMigrateUserModel(db)
	repositary := repositary.NewUserRepository(db)
	jwtManager := &authentication.JwtManager{SecretKey: os.Getenv("JWT_SECRET")}
	serv := service.NewUserService(repositary, jwtManager)
	api := api.NewApi(serv)

	r.POST("/checkuser", api.CheckUserApi)
	r.POST("/register", api.RegisterApi)
	r.POST("/login", api.LoginApi)

}
