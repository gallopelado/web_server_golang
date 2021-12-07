package main

import (
	"net/http"
)

type Router struct {
	//mapa anterior
	//rules map[string]http.HandlerFunc
	//mapa de mapas
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// mapa de reglas
func (r *Router) Findhandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf(w, "Hello world")
	handler, exist, methodExist := r.Findhandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
