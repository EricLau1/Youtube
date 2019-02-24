package routes

import(
  "net/http"
  "github.com/gorilla/mux"
  "go-webapp/middleware"
)

func NewRouter() *mux.Router {
  r := mux.NewRouter()
  r.HandleFunc("/", homeGetHandler).Methods("GET")
  r.HandleFunc("/", homePostHandler).Methods("POST")
  r.HandleFunc("/register", registerGetHandler).Methods("GET")
  r.HandleFunc("/register", registerPostHandler).Methods("POST")
  r.HandleFunc("/login", loginGetHandler).Methods("GET")
  r.HandleFunc("/login", loginPostHandler).Methods("POST")
  r.HandleFunc("/logout", logoutGetHandler).Methods("GET")
  r.HandleFunc("/admin", middleware.AuthRequired(adminGetHandler)).Methods("GET")
  r.HandleFunc("/products", middleware.AuthRequired(productGetHandler)).Methods("GET")
  r.HandleFunc("/product-create", middleware.AuthRequired(productCreateGetHandler)).Methods("GET")
  r.HandleFunc("/product-create", middleware.AuthRequired(productCreatePostHandler)).Methods("POST")
  r.HandleFunc("/product-edit", middleware.AuthRequired(productEditGetHandler)).Methods("GET")
  r.HandleFunc("/product-edit", middleware.AuthRequired(productEditPostHandler)).Methods("POST")
  r.HandleFunc("/product-delete", middleware.AuthRequired(productDeleteGetHandler)).Methods("GET")
  r.HandleFunc("/users", middleware.AuthRequired(userGetHandler)).Methods("GET")

  fileServer := http.FileServer(http.Dir("./assets/"))
  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))
  return r
}
