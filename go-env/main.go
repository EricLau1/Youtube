package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal(err)
  }
  port := os.Getenv("APP_PORT")
  message := os.Getenv("SECRET_MESSAGE")
  if port == "" {
    port = ":3000"
  }
  fmt.Println("Running on port", port)
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(message))
  })
  log.Fatal(http.ListenAndServe(port, nil))
}

