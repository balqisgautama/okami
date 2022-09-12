package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"okami.auth.backend/config"
	"okami.auth.backend/router/endpoint"
	"strconv"
)

func ApiController(port int) {
	handler := mux.NewRouter()

	//// --------- registration
	//handler.HandleFunc(setPath("/registration"), endpoint.RegistrationEndpoint.RegistrationEndpoint)
	////---------- email
	handler.HandleFunc(setPath("/email/validate"), endpoint.EmailEndpoint.ValidateEmailEndpoint)
	handler.HandleFunc(setPath("/email/resend"), endpoint.EmailEndpoint.ResendEmailEndpoint)
	////---------- login
	//handler.HandleFunc(setPath("/login"), endpoint.LoginEndpoint.GeneralLoginEndpoint)
	////---------- token
	//handler.HandleFunc(setPath("/token/generate"), endpoint.TokenEndpoint.TokenAuthFirebase)
	//---------- health
	handler.HandleFunc(setPath("/auth/health"), endpoint.HealthEndpoint.HealthStatus)
	handler.Handle(setPath("/auth/health/prometheus"), promhttp.Handler()).Methods("GET", "OPTIONS")
	//---------- resource
	handler.HandleFunc(setPath("/auth/resource"), endpoint.ResourceEndpoint.CRUDResource)
	//---------- token
	handler.HandleFunc(setPath("/auth/token/resource"), endpoint.TokenEndpoint.TokenResource)
	//---------- user
	handler.HandleFunc(setPath("/auth/user"), endpoint.UserEndpoint.CRUDUser)
	//---------- pcke
	handler.HandleFunc(setPath("/auth/pkce/step1"), endpoint.PKCEEndpoint.PKCEStep1)
	handler.HandleFunc(setPath("/auth/pkce/step2"), endpoint.PKCEEndpoint.PKCEStep2)
	handler.HandleFunc(setPath("/auth/pkce/step3"), endpoint.PKCEEndpoint.PKCEStep3)

	handler.Use(Middleware)
	fmt.Println(http.ListenAndServe(":"+strconv.Itoa(port), handler))
}

func setPath(path string) string {
	prefixPath := config.ApplicationConfiguration.GetServerPrefixPath()
	return "/" + prefixPath + path
}
