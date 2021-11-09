package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RequestBody struct {
	Mail string `json:"mail"`
}

func CheckUser(c *gin.Context) {

	var jsonBody RequestBody

	c.BindJSON(&jsonBody)

	c.JSON(200, gin.H{
		"doesExist": true,
		"userName":  "dummy",
	})

	fmt.Println(jsonBody.Mail)

}
