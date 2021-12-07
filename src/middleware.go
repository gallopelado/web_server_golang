package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Simulando un chequeo de login
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Checking authentication")
			flag := true
			if flag {
				fmt.Println("Successfully authentication")
				// llamar al siguiente middleware
				f(w, r)
			} else {
				fmt.Println("Failed authentication")
				return
			}
		}
	}
}

func Loggin() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// defer nos permite ejecutar al final de todo
			// bonita función anónima
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			// llamar al siguiente middleware
			f(w, r)
		}
	}
}
