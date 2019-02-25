package utils

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
)

func BodyParser(r *http.Request) []byte {
  body, _ := ioutil.ReadAll(r.Body)
  return body
}

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
  w.Header().Set("Content-type", "application/json; charset=UTF8")
  w.WriteHeader(statusCode)
  err := json.NewEncoder(w).Encode(data)
  CheckError(err)
}

func CheckError(err error) {
  if err != nil {
    log.Fatal(err)
  }
}
