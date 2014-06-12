package rest

import (
	"encoding/json"
	"github.com/kidstuff/WebUtil/config"
	"github.com/kidstuff/WebUtil/response"
	"net/http"
	"strings"
)

func GetConf(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	cfg, err := config.Provider().OpenConfigurator(req)
	if err != nil {
		response.InternalErrorResponse(rw, &response.JSONErr{
			Message:     err.Error(),
			Description: "Error when open Configurator.",
		})
		return
	}
	defer cfg.Close()

	result := make(map[string]string)
	ks := strings.Split(req.FormValue("key"), ",")
	if len(ks) == 0 {
		response.BadRequestResponse(rw, &response.JSONErr{
			Message:     "Must have one more key in request parameter",
			Description: "",
		})
		return
	}

	for _, k := range ks {
		v, err := cfg.Get(k)
		if err != nil {
			result[k] = ""
		} else {
			result[k] = v
		}
	}

	json.NewEncoder(rw).Encode(&result)
}

func SetConf(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")

	cfg, err := config.Provider().OpenConfigurator(req)
	if err != nil {
		response.InternalErrorResponse(rw, &response.JSONErr{
			Message:     err.Error(),
			Description: "Error when open Configurator.",
		})
		return
	}
	defer cfg.Close()

	m := make(map[string]string)
	err = json.NewDecoder(req.Body).Decode(&m)
	if err != nil {
		response.BadRequestResponse(rw, &response.JSONErr{
			Message:     err.Error(),
			Description: "Error when unmarshal json data.",
		})
		return
	}
	defer req.Body.Close()

	for k, v := range m {
		cfg.Set(k, v)
	}
}
