package handler

import (
	"log"
	"net/http"
	"strings"

	store "go-api/store"
)

var db = store.New()

func GetHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		KeyId := strings.Split(r.URL.Path, "/")[3]

		Data := db.Get(KeyId)

		log.Println("Get Key = " + KeyId)
		log.Println("Get Value = " + Data)


		w.WriteHeader(http.StatusOK)
		w.Write([]byte(Data))
	}
}
func AddHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		key := r.FormValue("key")

		incrementId := db.Add(key)

		log.Println("Added Key = " + key)
		log.Println("Added Id = " + incrementId)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(incrementId))
	}
}

func FlushHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		log.Println("Flush handler started")

		db.Flush()

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Flush Congratulations"))
	}
}
