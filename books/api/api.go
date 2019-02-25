package api

import (
  "net/http"
  "log"
  "fmt"
  "books/api/routes"
  "books/api/models"
)

func Run() {
  db := models.Connect()
  defer db.Close()
  if !db.HasTable(&models.Book{}) {
    db.Debug().CreateTable(&models.Book{})
  }
  listen(3000)
}

func listen(p int) {
  port := fmt.Sprintf(":%d", p)
  fmt.Printf("\n\nListening port %s...\n", port)
  r := routes.NewRouter()
  log.Fatal(http.ListenAndServe(port, r))
}
