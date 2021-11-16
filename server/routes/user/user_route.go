package user

import (
	userdb "backend/server/routes/user/data/user_db"
	"fmt"

	"github.com/gin-gonic/gin"
)

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

func CheckUser(c *gin.Context) {

	var jsonBody RequestBody
	c.BindJSON(&jsonBody)

	doesUserExist := userdb.CheckIfUserExists(jsonBody.Mail)

	c.JSON(200, gin.H{
		"doesExist": doesUserExist,
		"userName":  "dummy",
	})

	fmt.Println(jsonBody.Mail)

}

func RegisterUser(c *gin.Context) {
	var jsonBody RegisterRequestBody
	c.BindJSON(&jsonBody)

	registerModel, err := userdb.RegisterUser(jsonBody.Mail, jsonBody.Username, jsonBody.Password)
	if err != nil {
		c.JSON(500, gin.H{"isRegistrationCompleted": false})
		return
	}
	c.JSON(200, registerModel)
}

func LoginUser(c *gin.Context) {
	var jsonBody LoginRequestBody
	c.BindJSON(&jsonBody)
	result := userdb.LoginUser(jsonBody.Mail, jsonBody.Password)
	fmt.Println(result)
	c.JSON(200, result)
}
