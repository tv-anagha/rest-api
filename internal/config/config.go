package config

import "github.com/spf13/viper"

type Config struct {
	Env string 
	StoragePath string 
	HTTPServer
}

type HTTPServer struct {
		Address string 
} 

func LoadConfig(path string) (*Config, error) {