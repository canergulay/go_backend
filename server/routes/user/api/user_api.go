package api

import (
	"backend/server/routes/user/model"
	"backend/server/routes/user/service"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	service *service.UserLoginRegisterService
}

func NewService(s *service.UserLoginRegisterService) *UserApi {
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
	registerModel, err := a.service.Register(requestBody.Mail, requestBody.Username, requestBody.Password)
	if err != nil {
		c.JSON(500, gin.H{"isRegistrationCompleted": false})
		return
	}
	c.JSON(200, registerModel)
}

func (a *UserApi) CheckUserApi(c *gin.Context) {
	var requestBody model.RequestBody
	c.BindJSON(&requestBody)
	result := a.service.CheckIfUserExists(requestBody.Mail)
	c.JSON(200, result)
}
