package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yarm:"env" env-required:"true"`
	StoragePath string `yarm:"storage_path" env-required:"true"`
	HTTPServer  `yarm:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to config file")

		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is empty")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatal("config file does not exist: " + configPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatal("connot read config: " + err.Error())
	}

	return &cfg

}
