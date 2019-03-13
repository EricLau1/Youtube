package controllers

import (
	"net/http"
	"go-api-products/api/utils"
)

type ApiInfo struct {
	Name string `json:"name"`
	Version string `json:"version"`
	CreatedAt string `json:"created_at"`
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, ApiInfo{"Api Crud Products", "1.0", "11 mar 2019 T20:02"}, http.StatusOK)
}