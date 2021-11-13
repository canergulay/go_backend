package server

import (
	landing "backend/server/routes/landing"
	searchcourse "backend/server/routes/search_course"
	user "backend/server/routes/user"
	"backend/server/routes/user/data/user_db"

	"github.com/gin-gonic/gin"
)

func Run() {
	user_db.AutoMigrate()
	//routes
	r := gin.Default()
	r.GET("/", landing.MainRouter)
	r.POST("/checkuser", user.CheckUser)
	r.POST("/register", user.RegisterUser)
	r.POST("/login", user.LoginUser)
	r.POST("/search", searchcourse.SearchCourse)
	r.Run(":80")
}
