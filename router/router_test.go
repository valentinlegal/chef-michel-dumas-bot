package router

import (
	"testing"

	"github.com/valentinlegal/chef-michel-dumas-bot/data"
)

func TestPickWithValidFile(t *testing.T) {
	_ = data.Load("../data/data.json")

	_, err := pickGif()
	if err != nil {
		t.Errorf("Pick should be worked with a filled array of GIFs")
	}

	_, err = pickQuote()
	if err != nil {
		t.Errorf("Pick should be worked with a filled array of quotes")
	}
}

func TestPickWithInvalidFile(t *testing.T) {
	_ = data.Load("")

	_, err := pickGif()
	if err != nil {
		t.Errorf("Error should be returned when using an empty array of GIFs")
	}

	_, err = pickQuote()
	if err != nil {
		t.Errorf("Error should be returned when using an empty array of quotes")
	}
}
