package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

func InitCovert() (*map[string]map[string]interface{}, error) {
	file, err := os.ReadFile("./config/covert.yaml")
	if err != nil {
		return nil, err
	}

	var data map[string]map[string]interface{}
	if err := yaml.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return &data, nil

}
