package data

import "testing"

func TestLoadWithValidFilename(t *testing.T) {
	err := Load("data.json")
	if err != nil {
		t.Errorf("Load should be worked with a valid filename")
	}
}

func TestLoadWithInvalidFilename(t *testing.T) {
	err := Load("")
	if err == nil {
		t.Errorf("Error should be returned when using an invalid filename")
	}
}
