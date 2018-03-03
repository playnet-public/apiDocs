package main

import (
	"io/ioutil"
	"testing"
)

func TestRenderHTML(t *testing.T) {
	fileContents, err := ioutil.ReadFile("test/someBlueprint.apib")
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(fileContents)

	_, err = RenderHTML(&RequestBody{
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
