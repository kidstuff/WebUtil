package response

import (
	"encoding/json"
	"net/http"
)

const (
	ErrCodeUnknown   float32 = 0.0
	ErrCodeNotLogged float32 = 0.1
	ErrNoPermission  float32 = 0.2
)

// JSONErr returned to the user an informative error json message
type JSONErr struct {
	Code        float32
	Message     string
	Description string    `json:",omitempty"`
	StackStrace string    `json:",omitempty"`
	Errors      []JSONErr `json:",omitempty"`
}

func ErrorResponse(rw http.ResponseWriter, sttCode int, err *JSONErr) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(sttCode)
	json.NewEncoder(rw).Encode(err)
}

func BadRequestResponse(rw http.ResponseWriter, err *JSONErr) {
	ErrorResponse(rw, http.StatusBadRequest, err)
}

func ForbiddenResponse(rw http.ResponseWriter, err *JSONErr) {
	ErrorResponse(rw, http.StatusForbidden, err)
}

func InternalErrorResponse(rw http.ResponseWriter, err *JSONErr) {
	ErrorResponse(rw, http.StatusInternalServerError, err)
}

func UnauthorizedResponse(rw http.ResponseWriter, err *JSONErr) {
	ErrorResponse(rw, http.StatusUnauthorized, err)
}
