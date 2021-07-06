package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views")))
	http.ListenAndServe(":8080", router)
}
