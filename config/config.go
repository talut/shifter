package config

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Name   string
	Port   string
	Routes []Route
)

type Config struct {
	Name   string  `json:"name"`
	Port   string  `json:"port"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Key  string `json:"key"`
	From From   `json:"from"`
	To   To     `json:"to"`
}

type From struct {
	Method string `json:"method"`
}

type To struct {
	Method string `json:"method"`
	Url    string `json:"url"`
}

func Parse(path *string) {
	// read file from file path
	data, err := os.ReadFile(*path)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	// unmarshal json to Config struct
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("Unable to unmarshal JSON: %v", err)
	}

	// set global variables
	Name = config.Name
	Port = config.Port
	Routes = config.Routes
}
