package util

import (
	"github.com/dgrijalva/jwt-go"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"strconv"
	"time"
)

type PayloadJWTUser struct {
	ClientID string `json:"client_id"`
	jwt.StandardClaims
}

// GenerateJWTUser ---------------------------------------------------------------
// created at 08-31-2022
func GenerateJWTUser(clientID string, id int64) (string, string) {
	return generateJWTUser(clientID, id)
}

// created at 08-31-2022
func generateJWTUser(clientID string, id int64) (string, string) {
	tokenCode := PayloadJWTResource{
		ClientID: clientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   strconv.Itoa(int(id)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetJWTToken().JWT))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// ValidateJWTUser ---------------------------------------------------------------
// created at 08-31-2022
func ValidateJWTUser(jwtToken string) (output *PayloadJWTUser, response string) {
	token, response := jwtValidatorUser(jwtToken)
	if response != "" {
		return nil, response
	}
	output = token.Claims.(*PayloadJWTUser)

	return output, response
}

// created at 08-31-2022
func jwtValidatorUser(jwtToken string) (*jwt.Token, string) {
	var token *jwt.Token
	var err error

	claims := &PayloadJWTUser{}
	token, err = jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ApplicationConfiguration.GetJWTToken().JWT), nil
	})

	if err != nil {
		return token, err.Error()
	}

	if token.Header["alg"] != "HS512" {
		return token, constanta.InvalidToken
	}
	return token, ""
}
