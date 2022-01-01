package server

import (
	"backend/global/authentication"
	"backend/global/constants"
	"backend/server/routes/groups"
	"backend/server/routes/message"
	searchcourse "backend/server/routes/search_course"
	"backend/server/routes/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	pgDB       *gorm.DB
	gin        *gin.Engine
	jwtManager *authentication.JwtManager
	logger     *log.Logger
}

func (app *App) KickASS(pg *gorm.DB, gin *gin.Engine, jwtManager *authentication.JwtManager) {
	app.pgDB = pg
	app.gin = gin
	app.jwtManager = jwtManager

	app.run()
}

func (app *App) run() {

	app.gin.POST("/search", searchcourse.SearchCourse)
	// FROM NOW ON, WE'LL CONTINUE WITH CLEAN ARCHITECTURE //
	// THE CODE LINE BELOW WILL REPRESENT A WHOLE PACKAGE
	//-- INITS ALL ENDPOINTS FOR GROUP ENDPOINT --//
	user.InitUserRouter(app.gin, app.pgDB, app.jwtManager)
	groups.InitGroupRouter(app.gin, app.pgDB)
	message.InitMessageRouter(app.gin, app.pgDB, app.jwtManager)
	staticPathConfigurations(app.gin)
	app.gin.Run(":80")
}

func staticPathConfigurations(gin *gin.Engine) {
	gin.Static(constants.GroupImagePath, fmt.Sprintf("../..%s", constants.GroupImagePath))
	gin.Static(constants.UserImagePath, fmt.Sprintf("../..%s", constants.UserImagePath))
}
