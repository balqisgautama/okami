package util

import (
	"github.com/dgrijalva/jwt-go"
	"okami.auth.backend/constanta"
	"strconv"
	"time"
)

type PayloadJWTResource struct {
	ClientID string `json:"client_id"`
	jwt.StandardClaims
}

// GenerateJWTResource ---------------------------------------------------------------
// created at 08-18-2022
func GenerateJWTResource(key string, issuer string, clientID string, id int64) (string, string) {
	return generateJWTResource(key, issuer, clientID, id)
}

// created at 08-18-2022
func generateJWTResource(key string, issuer string, clientID string, id int64) (string, string) {
	tokenCode := PayloadJWTResource{
		ClientID: clientID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   strconv.Itoa(int(id)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// ValidateJWTResource ---------------------------------------------------------------
// created at 08-18-2022
func ValidateJWTResource(jwtToken string, key string) (output *PayloadJWTResource, response string) {
	token, response := jwtValidatorResource(jwtToken, key)
	if response != "" {
		return nil, response
	}
	output = token.Claims.(*PayloadJWTResource)

	return output, response
}

// created at 08-18-2022
func jwtValidatorResource(jwtToken string, key string) (*jwt.Token, string) {
	var token *jwt.Token
	var err error

	claims := &PayloadJWTResource{}
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
