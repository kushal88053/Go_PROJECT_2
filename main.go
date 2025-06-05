package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	routers.regirsterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))

}
