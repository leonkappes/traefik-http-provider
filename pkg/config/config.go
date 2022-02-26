package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port int
}

func ParseConfig() (*Config, error) {
	jsonFile, err := os.Open("config.json")

	if err != nil {
		log.Fatalf("Error reading config.json: %s", err)
		return nil, err
	}

	defer jsonFile.Close()

	byteData, _ := ioutil.ReadAll(jsonFile)

	var cfg Config

	if err = json.Unmarshal(byteData, &cfg); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &cfg, nil
}
