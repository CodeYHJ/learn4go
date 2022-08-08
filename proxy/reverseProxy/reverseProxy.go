package reverseProxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// ReverseProxy implement
func OriginReverseProxy(domain string) func(http.ResponseWriter, *http.Request) {
	urlParse, err := url.Parse(domain)
	if err != nil {
		log.Fatalf("parse url fail: %v", err)
	}
	//generate reverseProxy
	reverseProxy := httputil.NewSingleHostReverseProxy(urlParse)
	//modify request
	originDirector := reverseProxy.Director
	reverseProxy.Director = func(req *http.Request) {
		//origin
		originDirector(req)
		//custom
		req.Host = urlParse.Host
		req.URL.Host = urlParse.Host
		req.URL.Scheme = urlParse.Scheme
		req.URL.Path = "/s"
		req.URL.RawQuery = "wd=reverse"
	}
	return func(w http.ResponseWriter, req *http.Request) {
		reverseProxy.ServeHTTP(w, req)
	}
}
