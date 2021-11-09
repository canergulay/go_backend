package user

import (
	userdb "backend/server/routes/user/data/user_db"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Mail string `json:"mail"`
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
