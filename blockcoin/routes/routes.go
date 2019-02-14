package routes

import (
  "github.com/gorilla/mux"
  "go-api/controllers"
  "go-api/middlewares"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter().StrictSlash(true)
  r.HandleFunc("/", controllers.GetHome).Methods("GET")
  r.HandleFunc("/users", middlewares.IsAuth(controllers.GetUsers)).Methods("GET")
  r.HandleFunc("/users", controllers.PostUsers).Methods("POST")
  r.HandleFunc("/users/{uid}", middlewares.IsAuth(controllers.GetUser)).Methods("GET")
  r.HandleFunc("/users/{uid}", middlewares.IsAuth(controllers.PutUser)).Methods("PUT")
  r.HandleFunc("/users/{uid}", middlewares.IsAuth(controllers.DeleteUser)).Methods("DELETE")
  r.HandleFunc("/wallets", middlewares.IsAuth(controllers.GetWallets)).Methods("GET")
  r.HandleFunc("/wallets/{public_key}", middlewares.IsAuth(controllers.GetWallet)).Methods("GET")
  r.HandleFunc("/wallets/{public_key}", middlewares.IsAuth(controllers.PutWallet)).Methods("PUT")
  r.HandleFunc("/transactions", middlewares.IsAuth(controllers.GetTransactions)).Methods("GET")
  r.HandleFunc("/transactions/{public_key}", middlewares.IsAuth(controllers.PostTransaction)).Methods("POST")
  r.HandleFunc("/login", controllers.Login).Methods("POST")
  return r
}

