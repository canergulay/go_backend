package service

import (
	"backend/global/authentication"
	"backend/server/routes/user/model"
	"backend/server/routes/user/repositary"
	"fmt"
)

type UserLoginRegisterService struct {
	repositary *repositary.UserDbRepositary
	jwtManager *authentication.JwtManager
}

func NewUserService(r *repositary.UserDbRepositary, jwtManager *authentication.JwtManager) *UserLoginRegisterService {
	return &UserLoginRegisterService{
		repositary: r,
		jwtManager: jwtManager,
	}
}

func (service *UserLoginRegisterService) Login(mail string, password string) model.LoginResponseModel {
	return service.repositary.LoginUser(mail, password)
}

func (service *UserLoginRegisterService) CheckIfUserExists(mail string) bool {
	return service.repositary.CheckIfUserExists(mail)
}

func (service *UserLoginRegisterService) Register(mail string, username string, password string) (*RegisterResponseWithJWTLaoded, error) {

	registerRM, err := service.repositary.RegisterUser(mail, username, password)
	if err != nil {
		fmt.Println(err)

		return nil, err
	}

	// IF REGISTRATION PROCESS IS SUCCESSFUL, WE WILL GET THE TOKENS
	credentials, err := service.jwtManager.JwtSignUpCredentialsCreator(&registerRM.User)
	if err != nil {
		return nil, err
	}

	// WE WILL RETURN BOTH REGISTER RESPONSE AND TOKENS

	return &RegisterResponseWithJWTLaoded{
		RegisterResponse: registerRM,
		Credentials:      *credentials,
	}, nil

}

type RegisterResponseWithJWTLaoded struct {
	RegisterResponse model.RegisterResponseModel         `json:"registerResponse"`
	Credentials      authentication.JwtSignUpCredentials `json:"tokens"`
}
