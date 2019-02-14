package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "go-api/routes"
  "go-api/models"
)

func main() {
  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }
  models.TestConnection()
  fmt.Printf("Api running on port %s\n", port)
  r := routes.NewRouter()
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
