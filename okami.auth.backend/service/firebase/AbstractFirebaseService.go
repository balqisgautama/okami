package service

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"okami.auth.backend/constanta"
	out "okami.auth.backend/dto/out"
	"okami.auth.backend/service"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type firebaseService struct {
	FileName string
	service.AbstractService
}

var (
	funcName        string
	FirebaseService = firebaseService{FileName: "AbstractFirebaseService.go"}
)

// created at 07-31-2022
func (input firebaseService) generalSetting() *auth.Client {
	funcName = "generalSetting"
	fmt.Println(funcName)
	opt := option.WithCredentialsFile(os.Getenv("OkamiConfiguration") + "firebase\\serviceAccountKey.json")
	fmt.Println("opt", opt)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	fmt.Println("err 1", err.Error())
	client, err := app.Firestore(context.Background())
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	fmt.Println("err 2", err.Error())

	auth, err := app.Auth(context.Background())
	if err != nil {
		fmt.Errorf("error initializing app: %v", err)
	}
	fmt.Println("err 3", err.Error())

	defer client.Close()
	return auth
}

// created at 07-31-2022
func (input firebaseService) generateCustomToken() (token string, output out.Payload) {
	funcName = "generateCustomToken"
	fmt.Println(funcName)
	token, err := input.generalSetting().CustomToken(context.Background(), "firebase_UID")
	if err != nil {
		output.Status.Code = constanta.CodeAuthorizationFailed
		output.Status.Message = err.Error()
		output.Status.Detail = funcName
		return
	}
	return
}
