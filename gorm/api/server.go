package api

import (
  "net/http"
  "fmt"
  "log"
  "gorm/api/routes"
)

func listen(p int) {
  fmt.Printf("\n\nListening port %d...", p)
  port := fmt.Sprintf(":%d", p)
  r := routes.NewRouter()
  log.Fatal(http.ListenAndServe(port, r))
}
