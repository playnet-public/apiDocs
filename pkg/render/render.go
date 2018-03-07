package render

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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
	DefaultTemplate string
	Engine          snowboard.Parser
}

func NewEndpoints(engine snowboard.Parser, template string) *Endpoints {
	return &Endpoints{
		Engine:          engine,
		DefaultTemplate: template,
	}
}

func (e *Endpoints) RenderIt(w http.ResponseWriter, r *http.Request) {
	request, err := e.getRequestBody(r)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	if request.Template == "" {
		request.Template = e.DefaultTemplate
	}

	var res string
	if request.Action == "html" {
		res, err = e.renderHTML(request)
	} else if request.Action == "json" {
		res, err = e.renderJSON(request)
	}

	if err != nil {
		w.WriteHeader(400)
		return
	}

	_, err = w.Write([]byte(res))
	if err != nil {
		w.WriteHeader(400)
	}
}

func (e *Endpoints) renderHTML(request *RequestBody) (string, error) {
	reader := bytes.NewReader([]byte(request.Input))
	bp, err := snowboard.Parse(reader, e.Engine)

	if err != nil {
		return "", err
	}

	tfStream, _ := ioutil.ReadFile(request.Template)

	var buf bytes.Buffer
	err = render.HTML(string(tfStream), &buf, bp)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (e *Endpoints) renderJSON(request *RequestBody) (string, error) {
	result, err := snowboard.ParseAsJSON(bytes.NewReader([]byte(request.Input)), e.Engine)
	return string(result), err
}

func (e *Endpoints) getRequestBody(r *http.Request) (*RequestBody, error) {
	var request RequestBody
	err := json.NewDecoder(r.Body).Decode(&request)
	return &request, err
}
