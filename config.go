package main

import (
	"github.com/json-m/udm-api"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Config - some imported/exported to config file
type Config struct {
	User         string `yaml:"user"`
	Pass         string `yaml:"pass"`
	Host         string `yaml:"host"`
	Site         string `yaml:"site"`
	SkipInsecure bool   `yaml:"skip_insecure"`
	Client       *udm_api.Client
}

// init config
var config Config

func readConfig() error {
	// open config file
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal("readConfig.open:", err)
	}

	// read config file
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("readConfig.Decode:", err)
	}

	f.Close()
	return nil
}
