package gorestmath

import (
	"fmt"
	"net/http"
	"strings"
)

const ParseError string = "Couldn't parse requested operation"

func useFmt() {
	fmt.Println("argh")
}

func DoSomeMath(w http.ResponseWriter, r *http.Request) {
	// TODO: handle null URL

	path := r.URL.Path
	pathArray := strings.Split(path, "/")

	if len(pathArray) != 4 {
		returnError(ParseError, w)
		return
	}
	w.Write([]byte(`hi`))
}

func returnError(content string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(content))
}
