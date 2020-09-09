package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func (router *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := router.rules[path] // Verifica si existe el path.
	handler, methodExist := router.rules[path][method]
	return handler, methodExist, exist
}

func (router *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := router.FindHandler(request.URL.Path, request.Method)

	if !exist {
		writer.WriteHeader(http.StatusNotFound)
		// Responde al client.
		return
	}

	if !methodExist {
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}

	handler(writer, request)
}

func NewRouter() *Router {
	return &Router {
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}