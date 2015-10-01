package gorestmath

import (
	"fmt"
	"net/http"
	"strconv"
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
	var op func(int, int) int

	firstNum, err := strconv.Atoi(pathArray[2])
	if err != nil {
		returnError("Can't parse first digit", w)
		return
	}

	secondNum, err := strconv.Atoi(pathArray[3])
	if err != nil {
		returnError("Can't parse second digit", w)
		return
	}

	switch pathArray[1] {
	case "add":
		op = func(a, b int) int { return a + b }
	case "subtract":
		op = func(a, b int) int { return a - b }
	case "multiply":
		op = func(a, b int) int { return a * b }
	case "divide":
		if secondNum == 0 {
			returnError("Can't divide by 0", w)
			return
		}
		op = func(a, b int) int { return a / b }
	default:
		returnError("Bad op specified", w)
		return
	}

	result := op(firstNum, secondNum)

	resultJson := fmt.Sprintf("{'result':'%v'}", result)

	w.Write([]byte(resultJson))
}

func returnError(content string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(content))
}
