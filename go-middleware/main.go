package main

import (
  "fmt"
  "log"
  "time"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  _, err := w.Write([]byte("Hello World!"))
  if err != nil {
    log.Fatal(err)
  }
}

func exampleMiddleware(handler http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    log.Println(r.Method, r.Proto, r.Host, r.URL, time.Since(time.Now()))
    handler.ServeHTTP(w, r)
  })
}

func main() {
  fmt.Println("Running...")
  http.HandleFunc("/", exampleMiddleware(handler))
  http.Handle("/favicon.ico", http.NotFoundHandler())
  log.Fatal(http.ListenAndServe(":3000", nil))
}
