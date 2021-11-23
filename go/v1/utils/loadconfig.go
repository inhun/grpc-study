package utils

import (
	"io/ioutil"
	"encoding/json"
)


type Config struct {
	Url string `json:"url"`
}


func LoadConfig(filePath string) (*Config, error) {
	cfg := &Config{}

	dataBytes, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		return cfg, err
	}

	json.Unmarshal(dataBytes, cfg)

	return cfg, nil
}