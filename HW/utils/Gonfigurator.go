package utils

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	AdminToken   string
	Host         string
	Root         string
	DatabasePath string
}

var DefaultConfig = Config{
	AdminToken:   "admin",
	Host:         "127.0.0.1:8080",
	Root:         "",
	DatabasePath: "hw.db",
}

var defConfigLocation = "./hw.conf"

func ReadConfig() Config {
	var configLocation string
	if len(os.Args) > 1 {
		configLocation = os.Args[1]
	} else {
		configLocation = defConfigLocation
	}
	configBytes, err := ioutil.ReadFile(configLocation)
	if err != nil {
		log.Println("ERROR: Failed to read config file: " + err.Error())
		log.Println("Warning: use default config instead")
		return DefaultConfig
	}
	tomlData := string(configBytes)
	var config = DefaultConfig
	if _, err := toml.Decode(tomlData, &config); err != nil {
		log.Println("ERROR: Failed to parse config file: " + err.Error())
		log.Println("Warning: use default config instead")
		return DefaultConfig
	}
	log.Println("Root: " + config.Root)
	log.Println("Host: " + config.Host)
	log.Println("Database: " + config.DatabasePath)
	log.Println("Token: " + config.AdminToken)
	return config
}
