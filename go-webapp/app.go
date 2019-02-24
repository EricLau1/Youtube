package main

import(
  "fmt"
  "log"
  "os"
  "net/http"
  "go-webapp/routes"
  "go-webapp/utils"
  "go-webapp/models"
  "go-webapp/sessions"
)


func main() {
  models.TestConnection()
  sessions.SessionOptions("localhost", "/", 3600, true)
  port := os.Getenv("PORT")
  if port == "" {
    port = "8080"
  }
  fmt.Printf("Listening Port %s\n", port)
  utils.LoadTemplates("views/*.html")
  r := routes.NewRouter()
  http.Handle("/", r)
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

