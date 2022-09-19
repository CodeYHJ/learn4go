package main

import (
	"encoding/json"
	"fmt"
	"learn4go/bind/router"
	"net/http"
)

type TestData struct {
	A string `json:"a" query:"a" form:"a"`
	B string `json:"b" query:"b" form:"b"`
}

type Data struct {
	A string     `json:"a" query:"a" form:"a"`
	B string     `json:"b" query:"b" form:"b"`
	C TestData   `json:"c" form:"c"`
	D []string   `json:"d"`
	E []TestData `json:"e"`
}

func main() {

	api := router.New()
	api.Group("/get")
	api.Get("/query", func(writer http.ResponseWriter, request *http.Request) {
		test := request.URL.Query()
		a := test.Get("a")
		b := test.Get("b")
		fmt.Printf("%v:%v", a, b)
		r := fmt.Sprintf("Query: a:%v, b:%v", a, b)
		writer.Write([]byte(r))
	})

	api.Get("/query/bind", func(writer http.ResponseWriter, request *http.Request) {
		var data Data
		Bind(request, &data)

		r := fmt.Sprintf("BindQuery: a:%v, b:%v", data.A, data.B)
		writer.Write([]byte(r))
	})
	api.Group("/post")
	api.Post("/form", func(writer http.ResponseWriter, request *http.Request) {
		err := request.ParseMultipartForm(32 << 20)
		if err != nil {
			fmt.Printf("ParseForm err: %v", err)
			return
		}
		a := request.Form.Get("a")
		b := request.Form.Get("b")
		r := fmt.Sprintf("Form: a:%v, b:%v", a, b)

		writer.Write([]byte(r))
	})
	api.Post("/form/bind", func(writer http.ResponseWriter, request *http.Request) {
		var data Data
		Bind(request, &data)

		r := fmt.Sprintf("BindQuery: a:%v, b:%v", data.A, data.B)
		writer.Write([]byte(r))
	})
	api.Post("/urlencoded", func(writer http.ResponseWriter, request *http.Request) {
		request.ParseForm()

		a := request.Form.Get("a")
		b := request.Form.Get("b")
		r := fmt.Sprintf("Form: a:%v, b:%v", a, b)

		writer.Write([]byte(r))
	})
	api.Post("/urlencoded/bind", func(writer http.ResponseWriter, request *http.Request) {
		var data Data
		Bind(request, &data)

		r := fmt.Sprintf("BindQuery: a:%v, b:%v", data.A, data.B)
		writer.Write([]byte(r))
	})

	api.Post("/json", func(writer http.ResponseWriter, request *http.Request) {
		jsonDecoder := json.NewDecoder(request.Body)
		jsonDecoder.DisallowUnknownFields()
		var d Data
		err := jsonDecoder.Decode(&d)
		r := ""
		if err != nil {
			r = fmt.Sprintf("json parse error: %v", err)
			writer.Write([]byte(r))

		}
		r = fmt.Sprintf("Json: a:%v, b:%v", d.A, d.B)
		writer.Write([]byte(r))
	})
	api.Post("/json/bind", func(writer http.ResponseWriter, request *http.Request) {
		var d Data
		err := Bind(request, &d)
		if err != nil {
			writer.Write([]byte(fmt.Sprintf("json parse error: %v", err)))
		}
		writer.Write([]byte(fmt.Sprintf("Json: a:%v, b:%v,c:%v,d:%v;e:%v", d.A, d.B, d.C, d.D, d.E)))
	})

	err := http.ListenAndServe(":1323", nil)
	fmt.Printf("start process")

	if err != nil {
		fmt.Printf("http error: %v", err)
	}
	fmt.Printf("start success")
}
