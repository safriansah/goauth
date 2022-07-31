package controllers

import (
	"encoding/json"
	"goauth/models"
	"goauth/services"
	"net/http"
)

type Response struct {
	Data  []models.User `json:"data"`
	Code  int           `json:"message"`
	Error string
}

func GetUsers(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var db = services.GetDb()
	var users = models.GetUsers()
	var response Response

	err := db.Find(&users).Error
	if err != nil {
		response.Code = 500
		response.Error = err.Error()
	} else {
		response.Code = 200
		response.Data = users
	}

	json.NewEncoder(res).Encode(&response)
}
