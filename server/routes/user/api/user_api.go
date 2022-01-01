package api

import (
	"backend/server/routes/user/model"
	"backend/server/routes/user/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	service *service.UserLoginRegisterService
}

func NewApi(s *service.UserLoginRegisterService) *UserApi {
	return &UserApi{
		service: s,
	}
}

func (a *UserApi) LoginApi(c *gin.Context) {
	var requestBody model.LoginRequestBody
	c.BindJSON(&requestBody)
	model := a.service.Login(requestBody.Mail, requestBody.Password)
	c.JSON(200, model)
}

func (a *UserApi) RegisterApi(c *gin.Context) {
	var requestBody model.RegisterRequestBody
	c.BindJSON(&requestBody)
	registerModelJWTLoaded, err := a.service.Register(requestBody.Mail, requestBody.Username, requestBody.Password)
	fmt.Println(registerModelJWTLoaded)
	if err != nil {
		c.JSON(500, gin.H{"isRegistrationCompleted": false})
		return
	}

	c.JSON(200, registerModelJWTLoaded)
}

func (a *UserApi) CheckUserApi(c *gin.Context) {
	var requestBody model.RequestBody
	c.BindJSON(&requestBody)
	result := a.service.CheckIfUserExists(requestBody.Mail)
	m := make(map[string]bool)
	m["doesExist"] = result
	c.JSON(200, m)
}
