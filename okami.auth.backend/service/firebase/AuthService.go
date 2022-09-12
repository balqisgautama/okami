package service

//
//import (
//	"fmt"
//	"net/http"
//	"okami.auth.backend/constanta"
//	res "okami.auth.backend/dto/out"
//	"okami.auth.backend/model"
//)
//
//// created at 07-31-2022
//func (input firebaseService) GenerateToken(request *http.Request, context *model.ContextModel) (output res.Payload, header map[string]string) {
//	fmt.Println("GenerateToken")
//	token, output := input.generateCustomToken()
//
//	if output.Status.Code != "" {
//		fmt.Println("masuk sini")
//		return
//	}
//	fmt.Println("nggak masuk sana")
//
//	output.Status.Code = constanta.PayloadStatusCode
//	output.Data.Content = token
//	return
//}
