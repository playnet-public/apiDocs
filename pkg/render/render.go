package render

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/bukalapak/snowboard/render"

	snowboard "github.com/bukalapak/snowboard/parser"
)

type RequestBody struct {
	Action   string `json:"action,omitempty"`
	Template string `json:"template,omitempty"`
	Input    string `json:"input,omitempty"`
}

type Endpoints struct {
	Engine snowboard.Parser
}

func NewEndpoints(engine snowboard.Parser) *Endpoints {
	return &Endpoints{
		Engine: engine,
	}
}

func (e *Endpoints) RenderIt(w http.ResponseWriter, r *http.Request) {
	request, err := e.getRequestBody(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	var bytes []byte
	if request.Action == "html" {
		bytes, err = e.renderHTML(request)
	} else if request.Action == "json" {
		bytes, err = e.renderJSON(request)
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

func (e *Endpoints) renderHTML(request *RequestBody) ([]byte, error) {
	bp, err := snowboard.Load(request.Input, e.Engine)
	//bp, err := snowboard.Parse(bytes.NewReader(request.Input), e.Engine)

	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	err = render.HTML(request.Template, buf, bp)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (e *Endpoints) renderJSON(request *RequestBody) ([]byte, error) {
	return nil, nil //snowboard.ParseAsJSON(bytes.NewReader(request.Input), e.Engine)
}

func (e *Endpoints) getRequestBody(r *http.Request) (*RequestBody, error) {
	var request *RequestBody
	return request, json.NewDecoder(r.Body).Decode(request)
}
