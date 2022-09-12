package util

import (
	"okami.auth.backend/constanta"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type PayloadJWTInternal struct {
	Locale   string `json:"locale"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Resource string `json:"resource"`
	Version  string `json:"version"`
	jwt.StandardClaims
}

type PayloadJWTResource struct {
	ClientID string `json:"username"`
	jwt.StandardClaims
}

// created at 06-21-2022
func generateJWTActivation(key string, issuer string, version string, username string, email string, code int64) (string, string) {
	tokenCode := PayloadJWTInternal{
		Username: username,
		Email:    email,
		Version:  version,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   strconv.Itoa(int(code)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// created at 06-21-2022
func GenerateJWTActivation(key string, issuer string, version string, username string, email string, code int64) (string, string) {
	return generateJWTActivation(key, issuer, version, username, email, code)
}

// created at 06-21-2022
func jwtValidator(jwtToken string, key string) (*jwt.Token, string) {
	var token *jwt.Token
	var err error

	claims := &PayloadJWTInternal{}
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

// created at 06-21-2022
func ValidateJWT(jwtToken string, key string) (output *PayloadJWTInternal, response string) {
	token, response := jwtValidator(jwtToken, key)
	if response != "" {
		return nil, response
	}
	output = token.Claims.(*PayloadJWTInternal)

	return output, response
}

// created at 06-21-2022
//func ConvertTokenData(input interface{}) *PayloadJWTInternal {
//	bolB, _ := json.Marshal(input)
//	tokenData := PayloadJWTInternal{}
//	json.Unmarshal(bolB, &tokenData)
//	return &tokenData
//}

// created at 07-05-2022
func generateJWT(key string, issuer string, version string, username string, email string, userid int64) (string, string) {
	tokenCode := PayloadJWTInternal{
		Username: username,
		Email:    email,
		Version:  version,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer,
			Subject:   strconv.Itoa(int(userid)),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(key))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// created at 07-05-2022
func GenerateJWT(key string, issuer string, version string, username string, email string, code int64) (string, string) {
	return generateJWT(key, issuer, version, username, email, code)
}

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
