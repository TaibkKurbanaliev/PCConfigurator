package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strings"
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
	if !strings.Contains(key, ".") {
		_, ok := c.data[key]
		if !ok {
			return "", errors.New("key not found")
		}

		return c.data[key].(string), nil
	}

	keys := strings.Split(key, ".")
	result, ok := c.data[keys[0]]
	if !ok {
		return "", errors.New("key not found")
	}

	for i := 1; i < len(keys); i++ {
		value, ok := result.(map[string]interface{})
		if !ok {
			break
		}

		result, ok = value[keys[i]]
		if !ok {
			return "", errors.New("key not found")
		}
	}

	return result.(string), nil
}
