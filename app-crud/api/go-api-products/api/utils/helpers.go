package utils

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func BodyParser(r *http.Request) []byte {
	body, _ := ioutil.ReadAll(r.Body)
	return body
}

func ToJson(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	CheckErr(err)
}