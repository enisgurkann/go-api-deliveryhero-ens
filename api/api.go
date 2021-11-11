package api

import (
	"log"
	"net/http"

	. "go-api/handler"
)

func Start() {
	apiRoot := "/api"
	http.HandleFunc(apiRoot+"/add", AddHandler)
	http.HandleFunc(apiRoot+"/get/", GetHandler)
	http.HandleFunc(apiRoot+"/flush", FlushHandler)

	http.ListenAndServe(":8888", nil)
	log.Println("Api Listen :8888")

}
