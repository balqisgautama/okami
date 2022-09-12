package util

//
//// created at 06-21-2022
//func ValidateJWT(jwtToken string, key string) (output *PayloadJWTActivation, response string) {
//	token, response := jwtValidator(jwtToken, key)
//	if response != "" {
//		return nil, response
//	}
//	output = token.Claims.(*PayloadJWTActivation)
//
//	return output, response
//}
//
//// created at 06-21-2022
//func jwtValidator(jwtToken string, key string) (*jwt.Token, string) {
//	var token *jwt.Token
//	var err error
//
//	claims := &PayloadJWTActivation{}
//	token, err = jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
//		return []byte(key), nil
//	})
//
//	if err != nil {
//		return token, err.Error()
//	}
//
//	if token.Header["alg"] != "HS512" {
//		return token, constanta.InvalidToken
//	}
//	return token, ""
//}

// created at 06-21-2022
//func ConvertTokenData(input interface{}) *PayloadJWTActivation {
//	bolB, _ := json.Marshal(input)
//	tokenData := PayloadJWTActivation{}
//	json.Unmarshal(bolB, &tokenData)
//	return &tokenData
//}

//// created at 07-05-2022
//func generateJWT(key string, issuer string, version string, username string, email string, userid int64) (string, string) {
//	tokenCode := PayloadJWTActivation{
//		Username: username,
//		Email:    email,
//		Version:  version,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
//			IssuedAt:  time.Now().Unix(),
//			Issuer:    issuer,
//			Subject:   strconv.Itoa(int(userid)),
//		},
//	}
//
//	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, tokenCode)
//	token, err := jwtToken.SignedString([]byte(key))
//	if err != nil {
//		return "", err.Error()
//	}
//	return token, ""
//}
//
//// created at 07-05-2022
//func GenerateJWT(key string, issuer string, version string, username string, email string, code int64) (string, string) {
//	return generateJWT(key, issuer, version, username, email, code)
//}
