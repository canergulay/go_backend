package middlewares

import (
	"backend/global/authentication"

	"github.com/gin-gonic/gin"
)

var (
	noTokenFound  string = "Unauthorized, please enter a valid token !"
	tokenNotValid string = "Unauthorized, your token is not valid !"
)

func JwtVerifer(jmanager *authentication.JwtManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		token := c.Request.Header.Get("Authorization")

		// IF THERE IS NO A HEADER VALUE WITH THE KEY ABOVE, WE WILL SIMPLE TERMINATE THIS REQUEST AND RETURN 403
		if len(token) < 10 {
			c.JSON(403, noTokenFound)
			c.Abort()
		}
		//IF THERE IS A TOKEN, WE WILL THEN CHECK IF IT IS A VALID ONE
		user, err := jmanager.JwtCredentialsVerifier(token)
		if err != nil {
			c.JSON(402, tokenNotValid)
			c.Abort()
		}

		// OTHERWISE WE WILL SET OUR USER OBJECT THAT WE JUST PARSED FROM JWT PAYLOAD

		c.Set("user", user)
		c.Next()

	}
}
