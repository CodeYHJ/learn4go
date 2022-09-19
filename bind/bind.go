package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func Bind(request *http.Request, data interface{}) error {
	contentType := request.Header.Get("Content-Type")

	if contentType == "application/json" {
		b, err := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if err != nil {
		}

		err = json.Unmarshal(b, data)
		if err != nil {
			r := fmt.Sprintf("json parse error: %v", err)
			return errors.New(r)
		}
	}
	if contentType == "application/x-www-form-urlencoded" {
		err := request.ParseForm()
		if err != nil {
			r := fmt.Sprintf("ParseForm err: %v", err)
			return errors.New(r)
		}
		formData := request.PostForm

		bindData(formData, data, "form")
	}

	if strings.Contains(contentType, "multipart/form-data") {
		err := request.ParseMultipartForm(32 << 20)
		if err != nil {
			r := fmt.Sprintf("ParseForm err: %v", err)
			return errors.New(r)
		}
		formData := request.PostForm
		bindData(formData, data, "form")
	}

	if contentType == "" && request.Method == "GET" {
		queryData := request.URL.Query()
		bindData(queryData, data, "query")
	}

	return nil
}

func bindData(requestData map[string][]string, data interface{}, tag string) {
	typ := reflect.TypeOf(data).Elem()
	val := reflect.ValueOf(data).Elem()

	for i := 0; i < typ.NumField(); i++ {
		typeField := typ.Field(i)
		structField := val.Field(i)
		inputFieldName := typeField.Tag.Get(tag)
		inputValue, exists := requestData[inputFieldName]
		if !exists {
			fmt.Sprint("Map get fail")
		}
		structField.SetString(inputValue[0])
	}

}
