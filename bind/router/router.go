package router

import "net/http"

type Router struct {
	Prefix string
}

func New() *Router {
	return &Router{}
}

func (r *Router) Group(prefix string) {
	r.Prefix = prefix
}

func (r *Router) Get(path string, handler http.HandlerFunc) {
	http.HandleFunc(r.Prefix+path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) {
	http.HandleFunc(r.Prefix+path, handler)

}
