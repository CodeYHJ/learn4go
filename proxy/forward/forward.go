package forward

import (
	"io"
	"log"
	"net/http"
)

type Forward struct {
	Domain string
}

func (p *Forward) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	newReq, err := http.NewRequest("get", p.Domain, nil)
	if err != nil {
		log.Fatalf("Generate Request Err: %v", err)
	}
	client := &http.Client{}
	cRes, err := client.Do(newReq)
	if err != nil {
		log.Fatalf("Client Request Err: %v", err)
	}
	for key, value := range cRes.Header {
		for _, v := range value {
			rw.Header().Add(key, v)
		}
	}
	rw.WriteHeader(cRes.StatusCode)
	io.Copy(rw, cRes.Body)
	cRes.Body.Close()

}
