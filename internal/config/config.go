package config


import "os"
import "flag"
import "log"
import "github.com/ilyakaznacheev/cleanenv"

// env-default:"production"
// struct tags are used to map the environment variables to the struct fields
type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
		Address string `yaml:"address" env-required:"true"`
} 


func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")

		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatalf("config path is required")
		}
	} 

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("can not read config files:: %s", err.Error())
	}


	return &cfg
}
