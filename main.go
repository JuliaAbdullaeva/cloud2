//https://github.com/cloud-native-go/examples/tree/main/ch05
package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var store = make(map[string]string)
var ErrorNoSuchKey = errors.New("no such key")

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")
	log.Fatal(http.ListenAndServe(":9999", r))
}
