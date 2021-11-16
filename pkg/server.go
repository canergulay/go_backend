package server

import (
	searchcourse "backend/pkg/routes/search_course"
	user "backend/pkg/routes/user"
	"backend/pkg/routes/user/data/user_db"

	"github.com/gin-gonic/gin"
)

func Run() {
	user_db.AutoMigrate()
	//routes
	r := gin.Default()

	r.POST("/checkuser", user.CheckUser)
	r.POST("/register", user.RegisterUser)
	r.POST("/login", user.LoginUser)
	r.POST("/search", searchcourse.SearchCourse)
	r.Run(":80")
}
