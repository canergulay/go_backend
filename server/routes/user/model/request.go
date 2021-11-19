package model

type RequestBody struct {
	Mail string `json:"mail"`
}

type RegisterRequestBody struct {
	Mail     string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequestBody struct {
	Mail     string `json:"email"`
	Password string `json:"password"`
}
