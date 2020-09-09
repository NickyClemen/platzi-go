package main // Permite que los archivos en el mismo package lean los archivos.

import "net/http"

type Server struct {
	port string
	router *Router
}

func NewServer(port string) *Server {
	return &Server {
		port: port,
		router: NewRouter(),
	}
}

func (server *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := server.router.rules[path]

	if !exist {
		server.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	server.router.rules[path][method] = handler
}

func (server *Server) AddMiddleware(fn http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	// ... indica al compilador que se desconoce cantidad de parámetros.
	for _, middleware := range middlewares {
		fn = middleware(fn)
	}

	return fn
}

func (server *Server) Listen() error {
	http.Handle("/", server.router)

	err := http.ListenAndServe(server.port, nil)

	if err != nil {
		return err
	}

	return nil
}