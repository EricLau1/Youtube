package controllers

import (
  "net/http"
  "encoding/json"
  "strconv"
  "gorm/api/utils"
  "gorm/api/models"
  "github.com/gorilla/mux"
)

func PostUser(w http.ResponseWriter, r *http.Request) {
  body := utils.BodyParser(r)
  var user models.User
  err := json.Unmarshal(body, &user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  err = models.NewUser(user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, "Usuário adicionado com sucesso!", http.StatusCreated)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
  users := models.GetAll(models.USERS)
  utils.ToJson(w, users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  user := models.GetById(models.USERS, id)
  utils.ToJson(w, user, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 64)
  _, err := models.Delete(models.USERS, id)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return 
  }
  w.WriteHeader(http.StatusNoContent)
}

func PutUser(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.ParseUint(vars["id"], 10, 32)
  body := utils.BodyParser(r)
  var user models.User
  err := json.Unmarshal(body, &user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  user.Id = uint32(id)
  rows, err := models.UpdateUser(user)
  if err != nil {
    utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
    return
  }
  utils.ToJson(w, rows, http.StatusOK)
}
