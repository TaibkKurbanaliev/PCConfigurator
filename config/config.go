package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Config struct {
	data map[string]interface{}
}

func NewConfig(path string) *Config {
	if path == "" {
		return nil
	}

	configurationFile, err := readConfigurationFile(path)
	if err != nil {

		return nil
	}

	return &Config{data: configurationFile}
}

func readConfigurationFile(filePath string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var jsonData []byte
	jsonData, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		return nil, err
	}

	return jsonMap, nil
}

func (c *Config) Get(key string) (string, error) {
	if _, ok := c.data[key]; !ok {
		return "", errors.New("key not found")
	}

	return c.data[key].(string), nil
}
