package service

import (
	"backend/server/routes/user/model"
	"backend/server/routes/user/repositary"
)

type UserLoginRegisterService struct {
	repositary *repositary.UserDbRepositary
}

func NewUserService(r *repositary.UserDbRepositary) *UserLoginRegisterService {
	return &UserLoginRegisterService{
		repositary: r,
	}
}

func (service *UserLoginRegisterService) Login(mail string, password string) model.LoginResponseModel {
	return service.repositary.LoginUser(mail, password)
}

func (service *UserLoginRegisterService) CheckIfUserExists(mail string) bool {
	return service.repositary.CheckIfUserExists(mail)
}

func (service *UserLoginRegisterService) Register(mail string, username string, password string) (model.RegisterResponseModel, error) {
	return service.repositary.RegisterUser(mail, username, password)
}
