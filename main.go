package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", output)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func output(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "random number is :%v", randomnubers())
}

func randomnubers() int {
	return rand.Intn(1000)
}
