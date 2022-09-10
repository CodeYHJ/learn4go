package main

import (
	"fmt"
	"net/http"
)

type Group struct {
	Prefix string
}

func (g *Group) Get(path string, handler http.HandlerFunc) {
	http.HandleFunc(g.Prefix+path, handler)
}

type Router struct {
}

func (r *Router) Group(prefix string) *Group {
	return &Group{Prefix: prefix}
}

func main() {
	router := &Router{}
	g := router.Group("/test1")
	g.Get("/1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test1 success"))
	})

	o := router.Group("/test2")
	o.Get("/1", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("test2 success"))
	})
	err := http.ListenAndServe(":1323", nil)
	if err != nil {
		fmt.Printf("http error: %v", err)
	}
	fmt.Printf("start success")

}
