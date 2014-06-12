package rest

import (
	"github.com/gorilla/mux"
	"github.com/kidstuff/WebAuth/auth"
)

func Handler(router *mux.Router) {
	router.HandleFunc("/configurations", auth.OAuthHandleWrapper(GetConf, []string{"admin"}, nil)).Methods("GET")
	router.HandleFunc("/configurations", auth.OAuthHandleWrapper(SetConf, []string{"admin"}, nil)).Methods("POSTa")
}
