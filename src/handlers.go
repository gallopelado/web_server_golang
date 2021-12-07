package main

import (
	"fmt"
	"net/http"
)

// handler para la raíz "/"
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test home page")
}
