package main

import(
  "fmt"
  "net/http"
  "./routes"
)

func main() {

  fmt.Println("Listening port 3000")

  r := routes.NewRouter()

  http.ListenAndServe(":3000", r)

}
