package service

import (
	"okami.auth.backend/constanta"
	res "okami.auth.backend/dto/out"
	"okami.auth.backend/service"
	"okami.auth.backend/util"
)

type loginService struct {
	FileName string
	service.AbstractService
}

var LoginService = loginService{FileName: "LoginService.go"}

// created at 07-05-2022
func (input loginService) checkingPassword(passwordLogin string, passwordDB string) (output res.Payload) {
	check := util.CheckPasswordHash(passwordLogin, passwordDB)
	if !check {
		output.Status.Code = constanta.CodeLoginFailed
		output.Status.Message = constanta.InvalidPassword
		output.Status.Detail = "CheckPasswordHash"
		return
	}
	return
}
