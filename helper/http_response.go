package helper

import (
	"encoding/json"
	"net/http"
	"restfulapi/api"
)

func RespBadRequest(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Problem parsing request data."
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(api.RespBadRequest{
		Meta: &api.Meta{
			Code:    400,
			Message: msg,
		},
	})
	if err != nil {
		panic(err)
	}
}

func RespUnauthorized(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Unauthorized."
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	err := json.NewEncoder(w).Encode(api.RespUnauthorized{
		Meta: &api.Meta{
			Code:    401,
			Message: msg,
		},
	})
	if err != nil {
		panic(err)
	}
}

func RespNotFound(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Data is empty or not found."
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	err := json.NewEncoder(w).Encode(api.RespNotFound{
		Meta: &api.Meta{
			Code:    404,
			Message: msg,
		},
	})
	if err != nil {
		panic(err)
	}
}

func RespUnprocessableEntity(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Fields is unprocessable."
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)
	err := json.NewEncoder(w).Encode(api.RespUnprocessableEntity{
		Meta: &api.Meta{
			Code:    422,
			Message: msg,
		},
	})
	if err != nil {
		panic(err)
	}
}

func RespInternalServerError(w http.ResponseWriter, msg string) {
	if msg == "" {
		msg = "Internal server error."
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(api.RespInternalServerError{
		Meta: &api.Meta{
			Code:    500,
			Message: msg,
		},
	})
}
