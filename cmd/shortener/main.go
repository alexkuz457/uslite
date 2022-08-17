package main

import (
	"fmt"
	"net/http"
)

const addr = "localhost:8080"

func main() {

	http.HandleFunc("/", Handler)
	http.ListenAndServe(addr, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":

		loc := "http://localhost:8080/abc"

		//w.Header().Set("Location", loc) //http://localhost:8080"http://localhost:8080/"
		//w.WriteHeader(http.StatusTemporaryRedirect)
		//w.Write([]byte("res"))

		http.Redirect(w, r, loc, http.StatusTemporaryRedirect)

	default:

		fmt.Fprintf(w, r.Method+" - Метод не поддерживается")

	}

}
