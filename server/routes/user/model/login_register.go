package model

type LoginResponseModel struct {
	Status int
	User   User
}

type RegisterResponseModel struct {
	IsRegistrationCompleted bool
	User                    User
}
