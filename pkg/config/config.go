package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var cfg *config

type config struct {
	server struct {
		port int `yaml:"port"`
	} `yaml:"server"`
	client struct {
		serverAddress string `yaml:"address"`
	} `yaml:"client"`
}

func GetServerAddress() string {
	a := cfg.client.serverAddress
	if len(a) == 0 {
		return "localhost:9999"
	}
	return a
}

func LoadConfig(cfgFile string) error {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Could not read cfg file %w", err)
		//log.Panicf("Could not read cfg file %v", err)

	}

	log.Printf("Read Config file ... %s", viper.ConfigFileUsed())

	if err := viper.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("Could not unmarshall configuration data %w", err)
		//log.Panicf("Could not unmarshall configuration data %v", err)
	}

	return nil
}
