package util

import (
	"github.com/dgrijalva/jwt-go"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"time"
)

type PayloadPKCEStep1 struct {
	ClientID            string `json:"client_id"`
	Type                string `json:"type"`
	Scope               string `json:"scope"`
	CodeChallenge       string `json:"code_challenge"`
	CodeChallengeMethod string `json:"code_challenge_method"`
	jwt.StandardClaims
}

// GenerateJWTPCKEStep1 ---------------------------------------------------------------
// created at 08-26-2022
func GenerateJWTPCKEStep1(codeChallenge string) (string, string) {
	return generateJWTPCKEStep1(codeChallenge)
}

// created at 08-26-2022
func generateJWTPCKEStep1(codeChallenge string) (string, string) {
	tokenCode := PayloadPKCEStep1{
		Type:                constanta.AuthTypePKCE,
		CodeChallengeMethod: constanta.EncryptSHA256,
		CodeChallenge:       codeChallenge,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
	token, err := jwtToken.SignedString([]byte(config.ApplicationConfiguration.GetClientCredentialsClientID()))
	if err != nil {
		return "", err.Error()
	}
	return token, ""
}

// ValidateJWTPKCEStep1 ---------------------------------------------------------------
// created at 08-26-2022
func ValidateJWTPKCEStep1(jwtToken string) (output *PayloadPKCEStep1, response string) {
	token, response := jwtValidatorPKCEStep1(jwtToken, config.ApplicationConfiguration.GetClientCredentialsClientID())
	if response != "" {
		return nil, response
	}
	output = token.Claims.(*PayloadPKCEStep1)

	return output, response
}

// created at 08-26-2022
// updated at 08-30-2022
func jwtValidatorPKCEStep1(jwtToken string, key string) (*jwt.Token, string) {
	var token *jwt.Token
	var err error

	claims := &PayloadPKCEStep1{}
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
