package routes

import (
  "gorm/api/controllers"
  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  
  /* Users Routes */
  r.HandleFunc("/users", controllers.PostUser).Methods("POST")
  r.HandleFunc("/users", controllers.GetUsers).Methods("GET")
  r.HandleFunc("/users/{id}", controllers.GetUser).Methods("GET")
  r.HandleFunc("/users/{id}", controllers.PutUser).Methods("PUT")
  r.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
  
  /* Posts Routes */
  r.HandleFunc("/posts", controllers.PostPost).Methods("POST")
  r.HandleFunc("/posts", controllers.GetPosts).Methods("GET")
  r.HandleFunc("/posts/{id}", controllers.GetPost).Methods("GET")
  r.HandleFunc("/posts/{id}", controllers.PutPost).Methods("PUT")
  r.HandleFunc("/posts/{id}", controllers.DeletePost).Methods("DELETE")
  
  /* Feedbacks Routes */
  r.HandleFunc("/feedbacks", controllers.PostFeedback).Methods("POST")
  r.HandleFunc("/feedbacks", controllers.GetFeedbacks).Methods("GET")
  r.HandleFunc("/feedbacks/{id}", controllers.GetFeedback).Methods("GET")
  r.HandleFunc("/feedbacks/{id}", controllers.PutFeedback).Methods("PUT")
  r.HandleFunc("/feedbacks/{id}", controllers.DeleteFeedback).Methods("DELETE")
  return r
}
