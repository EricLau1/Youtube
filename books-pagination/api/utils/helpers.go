package utils

import (
  "net/http"
  "log"
  "encoding/json"
  "io/ioutil"
  "strconv"
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

func Pagination(r *http.Request, limit int) (int, int) {
  keys := r.URL.Query()
  if keys.Get("page") == "" {
    return 1, 0
  }
  page, _ := strconv.Atoi(keys.Get("page"))
  if page < 1 {
    return 1, 0
  }
  begin := (limit * page) - limit
  return page, begin
}
