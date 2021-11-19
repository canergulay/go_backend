package user

import (
	"backend/server/routes/user/api"
	"backend/server/routes/user/model"
	"backend/server/routes/user/repositary"
	"backend/server/routes/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(r *gin.Engine, db *gorm.DB) {
	model.AutoMigrateUserModel(db)
	repositary := repositary.NewUserRepository(db)
	serv := service.NewUserService(repositary)
	api := api.NewService(serv)

	r.POST("/checkuser", api.CheckUserApi)
	r.POST("/register", api.RegisterApi)
	r.POST("/login", api.LoginApi)

}
