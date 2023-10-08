package main

import (
	"gopkg.in/yaml.v3"
	UDM_API "jason.lv/UDM-API"
	"log"
	"os"
)

// Config - some imported/exported to config file
type Config struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Host   string `yaml:"host"`
	Site   string `yaml:"site"`
	Client UDM_API.Client
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
