package authentication

import (
	"backend/server/routes/user/model"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const tokenInfo string = "loginToken"

type JwtManager struct {
	SecretKey string
}

type Claimer struct {
	*jwt.StandardClaims
	TokenType string
	model.User
}
type JwtSignUpCredentials struct {
	AccessToken  string
	RefreshToken string
}

func (helper JwtManager) JwtSignUpCredentialsCreator(user *model.User) (*JwtSignUpCredentials, error) {
	accesTokenExpire, refreshTokenExpire := getExpires()
	t := jwt.New(jwt.SigningMethodHS256)
	accessToken, err := tokenCreator(t, accesTokenExpire, user, helper.SecretKey)
	if err != nil {
		fmt.Println(err)
		//TODO : IMPLEMENT YOUR LOG LIBRARY HERE, PROBABLY LOGRUS
		return nil, errors.New("an unexpected error has occured, so sorry")
	}
	refreshToken, err := tokenCreator(t, refreshTokenExpire, user, helper.SecretKey)
	if err != nil {
		fmt.Println(err)

		//TODO : IMPLEMENT YOUR LOG LIBRARY HERE, PROBABLY LOGRUS
		return nil, errors.New("an unexpected error has occured, so sorry")
	}
	credentials := &JwtSignUpCredentials{AccessToken: accessToken, RefreshToken: refreshToken}
	return credentials, nil
}

//Will return the user object if token is verified, otherwise will return an indicating error
func (helper JwtManager) JwtCredentialsVerifier(token string) (*model.User, error) {
	t, err := jwt.ParseWithClaims(token, &Claimer{}, func(t *jwt.Token) (interface{}, error) {
		return helper.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims := t.Claims.(*Claimer)
	return &claims.User, err
}

func tokenCreator(t *jwt.Token, expiresAt int64, user *model.User, SecretKey string) (string, error) {

	t.Claims = &Claimer{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		tokenInfo,
		*user,
	}
	return t.SignedString([]byte(SecretKey))
}

func getExpires() (int64, int64) {
	accesTE := time.Now().Add(time.Hour * 72).Unix()
	refreshTE := time.Now().Add(time.Hour * 1400).Unix()
	return accesTE, refreshTE
}
