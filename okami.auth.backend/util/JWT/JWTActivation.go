package util

import (
	"github.com/dgrijalva/jwt-go"
	"okami.auth.backend/constanta"
	"strconv"
	"time"
)

type PayloadJWTActivation struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWTActivation ------------------------------------------------------------------------------------------
// created at 06-21-2022
// updated at 08-24-2022
func GenerateJWTActivation(key string, username string, email string, userID int64, expiredAt int64) (string, string) {
	return generateJWTActivation(key, username, email, userID, expiredAt)
}

// created at 06-21-2022
// updated at 08-24-2022
func generateJWTActivation(key string, username string, email string, userID int64, expiredAt int64) (string, string) {
	tokenCode := PayloadJWTActivation{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt,
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(int(userID)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// ValidateJWTActivation ------------------------------------------------------------------------------------------
// created at 08-23-2022
func ValidateJWTActivation(jwtToken string, key string) (output *PayloadJWTActivation, response string) {
	//funcName := "ValidateJWTActivation"
	token, response := jwtValidator(jwtToken, key)
	if response != "" {
		return
	}
	output = token.Claims.(*PayloadJWTActivation)

	return
}

// created at 06-21-2022
func jwtValidator(jwtToken string, key string) (*jwt.Token, string) {
	var token *jwt.Token
	var err error

	claims := &PayloadJWTActivation{}
	token, err = jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return token, err.Error()
	}

	if token.Header["alg"] != "HS512" {
		return token, constanta.InvalidToken
	}
	return token, ""
}
