package routes

import (
	"go-api-products/api/controllers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.GetHome).Methods("GET")
	r.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	r.HandleFunc("/products/{id}", controllers.GetProduct).Methods("GET")
	r.HandleFunc("/products", controllers.PostProduct).Methods("POST")
	r.HandleFunc("/products/{id}", controllers.PutProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	return r
}