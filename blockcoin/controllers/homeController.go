package controllers

import (
  "net/http"
  "go-api/utils"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
  utils.ToJson(w, struct{
    Message string `json:"message"`
  }{
    Message: "Go RESTful Api",
  })
}
