package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Config represents the application configuration
(type Config struct {
	Database Database `json:"database"`
	Server   Server   `json:"server"`
})

// Database represents the database configuration
(type Database struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	User string `json:"user"`
	Password string `json:"password"`
})

// Server represents the server configuration
(type Server struct {
	Port int `json:"port"`
})

func InitConfig() *Config {
	config := &Config{}
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(configFile, config); err != nil {
		log.Fatal(err)
	}
	return config
}
