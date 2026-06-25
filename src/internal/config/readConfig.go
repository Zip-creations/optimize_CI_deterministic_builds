package config

import "os"
import "fmt"
import "encoding/json"

func ReadConfig(path string) (Config, error) {
	var config Config
	data, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf("Error while reading config file:\n %s\n %w", path, err)
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Error while unmarshalling config JSON:\n %w", err)
	}
	return config, nil
}
