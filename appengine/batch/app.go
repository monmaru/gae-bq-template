package batch

import (
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	r := mux.NewRouter()
	r.Path("/batch/bq/import").Handler(withAppHandler(importHandler)).Methods("GET")
	http.Handle("/", r)
}
