package user

import (
	"backend/global/authentication"
	"backend/server/routes/user/api"
	"backend/server/routes/user/model"
	"backend/server/routes/user/repositary"
	"backend/server/routes/user/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRouter(r *gin.Engine, db *gorm.DB, jM *authentication.JwtManager) {
	model.AutoMigrateUserModel(db)
	repositary := repositary.NewUserRepository(db)
	serv := service.NewUserService(repositary, jM)
	api := api.NewApi(serv)

	r.POST("/checkuser", api.CheckUserApi)
	r.POST("/register", api.RegisterApi)
	r.POST("/login", api.LoginApi)

}
