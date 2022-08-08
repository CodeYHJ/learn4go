package main

import (
	"learn4go/proxy/reverseProxy"
	"log"
	"net/http"
)

const ProxyDomain = "https://baidu.com"

func main() {
	//forward
	//f:=&forward.Forward{Domain: ProxyDomain}
	//http.Handle("/forward",f)
	//reverse
	http.HandleFunc("/reverse", reverseProxy.OriginReverseProxy(ProxyDomain))

	log.Fatal(http.ListenAndServe(":9090", nil))
}
