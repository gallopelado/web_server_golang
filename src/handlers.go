package main

import (
	"encoding/json"
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

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata Metadata
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v\n", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		//fmt.Println(user.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//fmt.Fprintf(w, "Payload %v\n", data)
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
