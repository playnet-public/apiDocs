package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bukalapak/snowboard/render"

	snowboard "github.com/bukalapak/snowboard/parser"
)

type RequestBody struct {
	Action   string `json:"action,omitempty"`
	Template string `json:"template,omitempty"`
	Input    []byte `json:"input,omitempty"`
}

func RenderIt(w http.ResponseWriter, r *http.Request) {
	request, err := getRequestBody(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var bytes []byte
	if request.Action == "html" {
		bytes, err = RenderHTML(request)
	} else if request.Action == "json" {
		bytes, err = RenderJSON(request)
	}

	if err != nil {
		w.WriteHeader(400)
		return
	}

	_, err = w.Write(bytes)
	if err != nil {
		w.WriteHeader(400)
	}
}

func RenderHTML(request *RequestBody) ([]byte, error) {
	fmt.Println("test 7")
	bp, err := snowboard.Parse(bytes.NewReader(request.Input), engine)

	fmt.Println("test 3")
	if err != nil {
		return nil, err
	}

	fmt.Println("test 4")
	buf := bytes.NewBuffer([]byte{})
	fmt.Println("test 5")
	err = render.HTML(request.Template, buf, bp)
	fmt.Println("test 6")
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func RenderJSON(request *RequestBody) ([]byte, error) {
	return snowboard.ParseAsJSON(bytes.NewReader(request.Input), engine)
}

func getRequestBody(r *http.Request) (*RequestBody, error) {
	var request *RequestBody
	return request, json.NewDecoder(r.Body).Decode(request)
}