package routes

import (
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "io/ioutil"
  "strconv"
  "fmt"
  "../models"
  "../utils"
)

func NewRouter() *mux.Router {

  r := mux.NewRouter().StrictSlash(true)
  
  r.HandleFunc("/users", usersGetHandler).Methods("GET")
  r.HandleFunc("/users", usersPostHandler).Methods("POST")
  r.HandleFunc("/users/{id}", userGetHandler).Methods("GET")
  r.HandleFunc("/users/{id}", userPutHandler).Methods("PUT")
  r.HandleFunc("/users/{id}", userDeleteHandler).Methods("DELETE")
  r.HandleFunc("/login", loginPostHandler).Methods("POST")

  return r;

}

func HttpInfo(r *http.Request) {

  fmt.Printf("%s/\t %s\t %s%s\t %s\n", r.Method, r.Proto, r.Host, r.URL, utils.GetDateTime())

}

func ContentTypeJson(w http.ResponseWriter) {

  w.Header().Set("Content-Type", "application/json")

}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {

  ContentTypeJson(w)

  HttpInfo(r)

  users, err := models.GetUsers()

  if err != nil {
  
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "BAD REQUEST",
      Status: 400,
    })
    return
  }
  
  json.NewEncoder(w).Encode(users)

}

func usersPostHandler(w http.ResponseWriter, r *http.Request) {
  
  ContentTypeJson(w)

  HttpInfo(r)
  
  body, _ := ioutil.ReadAll(r.Body)

  var user models.User

  err := json.Unmarshal(body, &user)

  _, err = models.NewUser(user)
 
  if err != nil {
  
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "UNPROCESSABLE ENTITY",
      Status: 422,
    })
    return
  }

  json.NewEncoder(w).Encode(struct{
    Message string
    Status int
  }{
    Message: "Created Successfully!",
    Status: 201,
  })

}


func userGetHandler(w http.ResponseWriter, r *http.Request) {
  
  ContentTypeJson(w)

  HttpInfo(r)

  params := mux.Vars(r)

  id, _ := strconv.Atoi(params["id"])

  user, err := models.GetUser(id)

  if err != nil {
  
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "BAD REQUEST",
      Status: 400,
    })
    return
  }

  json.NewEncoder(w).Encode(user)
}

func userPutHandler(w http.ResponseWriter, r *http.Request) {

  ContentTypeJson(w)

  HttpInfo(r)

  params := mux.Vars(r)

  id, _ := strconv.Atoi(params["id"])

  var user models.User
  
  body, _ := ioutil.ReadAll(r.Body)

  err := json.Unmarshal(body, &user)
 
  if err != nil {
  
    w.WriteHeader(http.StatusInternalServerError)

    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "UNPROCESSABLE ENTITY",
      Status: 422,
    })
    return
  }
  
  user.Id = id
  rows, err := models.UpdateUser(user)
  
  if err != nil {
  
    w.WriteHeader(http.StatusInternalServerError)

    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "UNPROCESSABLE ENTITY",
      Status: 422,
    })

    return
  }

  json.NewEncoder(w).Encode(struct{RowsAffected int64}{RowsAffected: rows,})

}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {

  ContentTypeJson(w)

  HttpInfo(r)

  params := mux.Vars(r)

  id, _ := strconv.Atoi(params["id"])

  rows, err := models.DeleteUser(id)

  if err != nil {
    
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "BAD REQUEST",
      Status: 400,
    })
    return
  }

  json.NewEncoder(w).Encode(struct{RowsAffected int64}{RowsAffected: rows,})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {

  ContentTypeJson(w)

  HttpInfo(r)

  body, _ := ioutil.ReadAll(r.Body)

  var user models.User

  err := json.Unmarshal(body, &user)

  if err != nil {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(struct{
      Error string
      Status int
    }{
      Error: "UNPROCESSABLE ENTITY",
      Status: 422,
    })
    return
  }
    
 
  if user.Email == "" || user.Password == "" {
      w.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(w).Encode(struct{
        Error string
        Status int
      }{
        Error: "UNAUTHORIZED",
        Status: 401,
      })
      return 
  }

  user, err = models.Signin(user.Email, user.Password)
  
  if err != nil {
      
      fmt.Println(err)
      w.WriteHeader(http.StatusInternalServerError)
      json.NewEncoder(w).Encode(struct{
        Error string
        Status int
      }{
        Error: "UNAUTHORIZED",
        Status: 401,
      })
      return 
  }

  
  json.NewEncoder(w).Encode(user)
  
}







