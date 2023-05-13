package environment

import (
	"encoding/json"
	"log"
	"os"
)

type DBConfig struct {
	Port     int    `json:"port"`
	Host     string `json:"host"`
	Username string `json:"username"`
	DBname   string `json:"dbname"`
	SSlmode  string `json:"sslmode"`
}

type Config struct {
	Port     int      `json:"port"`
	Host     string   `json:"host"`
	DBConfig DBConfig `json:"db-config"`
}

func NewConfig(configFile string) (*Config, error) {
	rawJSON, err := os.ReadFile(configFile)
	if err != nil {
		log.Printf("Can't read config file, err: %s\n", err.Error())
		return nil, err
	}

	var config Config
	err = json.Unmarshal(rawJSON, &config)
	if err != nil {
		log.Printf("Can't unmarshall config json, err: %s\n", err.Error())
		return nil, err
	}

	return &config, nil
}
