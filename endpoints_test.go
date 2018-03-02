package main

import (
	"io/ioutil"
	"testing"
)

func TestRenderHTML(t *testing.T) {
	t.Log("test 1")
	fileContents, _ := ioutil.ReadFile("test/someBlueprint.apib")

	t.Log("test 2")
	_, err := RenderHTML(&RequestBody{
		Action:   "html",
		Template: "",
		Input:    fileContents,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}

func TestRenderJSON(t *testing.T) {
	fileContents, _ := ioutil.ReadFile("test/someBlueprint.apib")

	_, err := RenderJSON(&RequestBody{
		Action:   "json",
		Template: "",
		Input:    fileContents,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
