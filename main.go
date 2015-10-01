package main

import (
	"gorestmatch"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", DoSomeMath)
	if err := http.ListenAndServe("0.0.0.0:7552", nil); err != nil {
		log.Fatal("Couldn't start server", err)
	}
}
