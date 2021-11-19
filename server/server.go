package server

import (
	"backend/global/constants"
	"backend/server/routes/groups"
	searchcourse "backend/server/routes/search_course"
	"backend/server/routes/user"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	pgDB *gorm.DB
	gin  *gin.Engine
}

func (app *App) KickASS(pg *gorm.DB, gin *gin.Engine) {
	app.pgDB = pg
	app.gin = gin
	app.Run()
}

func (app *App) Run() {

	app.gin.POST("/search", searchcourse.SearchCourse)
	// FROM NOW ON, WE'LL CONTINUE WITH CLEAN ARCHITECTURE //
	// THE CODE LINE BELOW WILL REPRESENT A WHOLE PACKAGE
	//-- INITS ALL ENDPOINTS FOR GROUP ENDPOINT --//
	user.InitUserRouter(app.gin, app.pgDB)
	groups.InitGroupRouter(app.gin, app.pgDB)
	app.gin.Static(constants.GroupImagePath, fmt.Sprintf("../..%s", constants.GroupImagePath))
	app.gin.Run(":80")
}
