package environment

import (
	"encoding/json"
	"os"

	log "github.com/sirupsen/logrus"
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
		log.WithError(err).Error("cannot read config file")
		return nil, err
	}

	var config Config
	err = json.Unmarshal(rawJSON, &config)
	if err != nil {
		log.WithError(err).Error("cannot unmarshall config json")
		return nil, err
	}
	return &config, nil
}
