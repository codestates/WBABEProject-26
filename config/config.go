package config

import (
	"flag"
	"os"

	"github.com/pelletier/go-toml/v2"
)

//Config
type Config struct {
	Server struct {
		Mode string
		Port int
	}		

	DB struct {
		Host string
		User string
		Pw string
		Name string
	}
}

//Config Load
func GetConfig() (*Config, error) {
	//커맨드에 들어온 값에 따라 처리
	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.toml", "Use config file")
	flag.Parse()

	c := new (Config)

	if file, err := os.Open(configPath); err != nil {
		return nil, err
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			return nil, err
		} else {
			return c, nil
		}
	}
}