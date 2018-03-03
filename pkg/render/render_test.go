package render

import (
	"io/ioutil"
	"testing"

	"github.com/bukalapak/snowboard/adapter/drafter"
)

var endpoints = NewEndpoints(&drafter.Parser{})

func TestRenderHTML(t *testing.T) {
	_, err := ioutil.ReadFile("test/someBlueprint.apib")
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = endpoints.renderHTML(&RequestBody{
		Action:   "html",
		Template: "",
		Input:    "test/someBlueprint.apib",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestRenderJSON(t *testing.T) {
	fileContents, _ := ioutil.ReadFile("test/someBlueprint.apib")

	_, err := endpoints.renderJSON(&RequestBody{
		Action:   "json",
		Template: "",
		Input:    "test/someBlueprint.apib",
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
