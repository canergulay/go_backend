package server

import (
	"backend/config/pg_manager"
	"backend/server/routes/groups"
	searchcourse "backend/server/routes/search_course"
	user "backend/server/routes/user"
	"backend/server/routes/user/data/user_db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	pgDB      *gorm.DB
	ginRouter *gin.Engine
}

func (a *App) KickASS(pg *gorm.DB, gin *gin.Engine) {
	a.pgDB = pg_manager.InitPostgreSQL()
	a.ginRouter = gin
}

func Run() {
	user_db.AutoMigrate()
	//routes
	r := gin.Default()

	r.POST("/checkuser", user.CheckUser)
	r.POST("/register", user.RegisterUser)
	r.POST("/login", user.LoginUser)
	r.POST("/search", searchcourse.SearchCourse)
	// FROM NOW ON, WE'LL CONTINUE WITH CLEAN ARCHITECTURE //
	// THE CODE LINE BELOW WILL REPRESENT A WHOLE PACKAGE
	//-- INITS ALL ENDPOINTS FOR GROUP ENDPOINT --//
	groups.InitGroupRouter(r, pg_manager.GetPostgresConnection())
	r.Run(":80")
}
