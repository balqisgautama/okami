package util

import (
	"bytes"
	"html/template"
	"math/rand"
	"okami.auth.backend/config"
	"okami.auth.backend/constanta"
	"okami.auth.backend/model"
	jwt "okami.auth.backend/util/JWT"
	"strconv"
	"time"
)

// created at 06-21-2022
// updated at 07-04-2022
//func GenerateActivation(email string, username string) (result model.ActivationAccount, output res.Payload) {
//	result.Code.Int64 = generateCode()
//	result.Status.Int64 = constanta.UserPending
//	result.Expire.Int64 = time.Now().Unix()
//	result.LinkValidate.String, output = generateLink(result.Code.Int64, email, username, "/email/validate")
//	result.LinkResend.String, output = generateLink(result.Code.Int64, email, username, "/email/resend")
//	result.EmailTo.String = email
//	result.Counter.Int64 = 0
//	return result, output
//}

// created at 06-21-2022
// updated at 08-22-2022
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	random := rand.Int63n(constanta.EmailCodeMaximumLength-constanta.EmailCodeMinimumLength+1) + constanta.EmailCodeMinimumLength
	//fmt.Println("random", random)
	return strconv.FormatInt(random, 10)
}

// created at 06-21-2022
// updated at 08-24-2022
func generateLink(path string, code string, user model.UserGeneral, expiredAt int64) (result string) {
	var EmailActivationLink = config.ApplicationConfiguration.GetServerHost() + ":" +
		strconv.Itoa(config.ApplicationConfiguration.GetServerPort()) + "/" +
		config.ApplicationConfiguration.GetServerPrefixPath() + path + "?"
	result = EmailActivationLink

	codeJWT, _ := jwt.GenerateJWTActivation(code, user.Username.String, user.Email.String,
		user.UserID.Int64, expiredAt)

	result += "code=" + codeJWT
	result += "&email_to=" + user.Email.String
	result += "&username=" + user.Username.String
	return result
}

//// created at 06-22-2022
//func GenerateHTMLGeneral(title string, h1 string, p string) (html string) {
//	html = "" +
//		"<html>" +
//		"<head>" +
//		"<title>" + title + "</title>" +
//		"</head>" +
//		"<body>" +
//		"<h1>" + h1 + "</<h1>" +
//		"<h5>" + p + "</h5>" +
//		"</body>" +
//		"</html>"
//	return
//}
//

type DataHTMLFile struct {
	Title        string
	BeforeButton string
	Button       string
	ButtonUrl    string
	AfterButton  string
}

// ParseHTMLFileToString ------------------------------------------------------------------------------------------
// created at 06-23-2022
func ParseHTMLFileToString(templateFileName string, data interface{}) (result string) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return constanta.ActivationFailed
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return constanta.ActivationFailed
	}
	result = buf.String()
	return result
}

// UserModelToActivationUserModel ------------------------------------------------------------------------------------------
// created at 08-22-2022
func UserModelToActivationUserModel(user model.UserGeneral) (result model.ActivationUser) {
	result.Counter.Int64 = 0
	code := generateCode()
	result.Code.String = code
	result.ExpiredAt.Time = time.Now().Add(constanta.EmailExpiredHour * time.Hour)
	result.Status.Int64 = constanta.UserPending
	result.UserID.Int64 = user.UserID.Int64
	result.EmailTo.String = user.Email.String
	result.EmailLinkValidate.String = generateLink("/email/validate", code, user,
		time.Now().Add(constanta.EmailExpiredHour*time.Hour).Unix())
	result.EmailLinkResend.String = generateLink("/email/resend", code, user, 0)
	return
}

// ToUpdatedActivationUserData ------------------------------------------------------------------------------------------
// created at 08-24-2022
func ToUpdatedActivationUserData(data model.ActivationUser, user model.UserGeneral) (result model.ActivationUser) {
	result.ActivationID.Int64 = data.ActivationID.Int64
	result.Counter.Int64 = data.Counter.Int64 + 1
	code := generateCode()
	result.Code.String = code
	result.ExpiredAt.Time = time.Now().Add(constanta.EmailExpiredHour * time.Hour)
	result.Status.Int64 = constanta.UserPending
	//result.UserID.Int64 = user.UserID.Int64
	result.EmailTo.String = user.Email.String
	result.EmailLinkValidate.String = generateLink("/email/validate", code, user,
		time.Now().Add(constanta.EmailExpiredHour*time.Hour).Unix())
	result.EmailLinkResend.String = generateLink("/email/resend", code, user, 0)
	return
}
