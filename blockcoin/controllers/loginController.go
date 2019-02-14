package controllers

import (
  "net/http"
  "io/ioutil"
  "go-api/auth"
  "go-api/utils"
  "go-api/models"
  "encoding/json"
)

func Login(w http.ResponseWriter, r *http.Request) {
  body, _ := ioutil.ReadAll(r.Body)
  var user models.User
  err := json.Unmarshal(body, &user)
  if err != nil {
    utils.ErrorResponse(w, err, http.StatusUnauthorized)
    return
  }
  userAuthenticate, err := auth.SignIn(user)
  if err != nil {
    utils.ErrorResponse(w, err, http.StatusUnauthorized)
    return
  }
  utils.ToJson(w, userAuthenticate)
}
