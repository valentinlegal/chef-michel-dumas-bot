package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Data struct {
	Activities []string  `json:"activities"`
	Commands   []Command `json:"commands"`
	Gifs       []string  `json:"gifs"`
	Quotes     []string  `json:"quotes"`
}

type Command struct {
	Keys        []string `json:"keys"`
	Description string   `json:"description"`
}

var data Data

func Load(filename string) error {
	log.Println("[INFO] Load data")

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(byteValue), &data)

	return nil
}

func Activities() []string {
	return data.Activities
}

func Commands() []Command {
	return data.Commands
}

func Gifs() []string {
	return data.Gifs
}

func Quotes() []string {
	return data.Quotes
}
