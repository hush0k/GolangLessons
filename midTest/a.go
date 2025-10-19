package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Dog struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type DogsHome struct {
	Name string `json:"name"`
	Dogs []Dog  `json:"dogs"`
}

func main() {
	http.HandleFunc("/home", logging(homeFunc))
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/db", logging(dbConnect))
	http.ListenAndServe(":8085", nil)

}

func homeFunc(w http.ResponseWriter, r *http.Request) {

	bang := Dog{
		Name: "Bang",
		Age:  3,
	}

	alabai := Dog{
		Name: "Alabai",
		Age:  1,
	}

	markus := Dog{
		Name: "Markus",
		Age:  4,
	}

	nanomi := DogsHome{
		Name: "Nanomi",
		Dogs: []Dog{bang, alabai, markus},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nanomi)
	fmt.Println("answer go")

}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "foo")
}

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}
