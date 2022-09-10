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

func (g *Group) File(path string, fileDir string) {
	f := http.Dir(fileDir)
	fs := http.FileServer(f)
	if len(g.Prefix) > 0 {
		http.Handle(g.Prefix+path, fs)

	} else {
		http.Handle(path, fs)
	}
}

func (g *Group) FilePrefix(path string, fileDir string) {
	var p string
	f := http.Dir(fileDir)
	fs := http.FileServer(f)
	if len(g.Prefix) > 0 {
		p = g.Prefix + path
	} else {
		p = path
	}
	h := http.StripPrefix(p, fs)
	http.Handle(path, h)
}

type Router struct {
}

func (r *Router) Group(prefix string) *Group {
	return &Group{Prefix: prefix}
}
func main() {
	router := &Router{}
	g := router.Group("")
	//only "/" path valid
	g.File("/", "./dist")

	// prefix path valid
	g.FilePrefix("/test/", "dist")

	err := http.ListenAndServe(":1323", nil)
	if err != nil {
		fmt.Printf("http error: %v", err)
	}
	fmt.Printf("start success")
}
