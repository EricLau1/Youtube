package routes

import (
  "books/api/controllers"
  "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/books", controllers.GetBooks).Methods("GET")
  r.HandleFunc("/books/{id}", controllers.GetBook).Methods("GET")
  r.HandleFunc("/books", controllers.PostBook).Methods("POST")
  r.HandleFunc("/books/{id}", controllers.PutBook).Methods("PUT")
  r.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
  return r
}
