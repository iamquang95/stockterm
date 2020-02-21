package config

import (
	"encoding/json"
	"io/ioutil"
)

func parseConfig(data []byte) (*Config, error) {
	config := &Config{}
	err := json.Unmarshal(data, config)
	return config, err
}

func GetConfig(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config, err := parseConfig(file)
	return config, err
}
