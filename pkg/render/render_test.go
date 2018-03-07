package render

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/bukalapak/snowboard/adapter/drafter"
)

var endpoints = NewEndpoints(&drafter.Engine{}, "../../test/templates/alpha.html")

func TestRenderIt(t *testing.T) {
	resp := httptest.NewRecorder()

	reader := bytes.NewReader([]byte(fmt.Sprintf("%v", RequestBody{
		Action:   "html",
		Template: "",
		Input:    "FORMAT: 1A\n\n# Tests\n\n# Group Test Blueprint\n## This is just a test url [/render]\n### There is only one method [GET]\n+ response 200 (application/html)\n",
	})))
	req := httptest.NewRequest("POST", "/render", reader)

	endpoints.RenderIt(resp, req)
}

func TestRenderHTML(t *testing.T) {
	file, err := ioutil.ReadFile("../../test/someBlueprint.apib")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = endpoints.renderHTML(&RequestBody{
		Action:   "html",
		Template: "../../test/templates/alpha.html",
		Input:    string(file),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestRenderJSON(t *testing.T) {
	file, err := ioutil.ReadFile("../../test/someBlueprint.apib")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = endpoints.renderJSON(&RequestBody{
		Action:   "json",
		Template: "",
		Input:    string(file),
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestGetRequestBody(t *testing.T) {
	fileCont, _ := ioutil.ReadFile("../../test/someBlueprintRequest.json")

	req := httptest.NewRequest("POST", "/render", bytes.NewReader(fileCont))
	_, err := endpoints.getRequestBody(req)
	if err != nil {
		t.Fatal(err.Error())
	}
}
