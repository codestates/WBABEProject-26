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
		DBName string
	}
}

//Config Load
func GetConfig() *Config {
	//커맨드에 들어온 값에 따라 처리
	var configPath string
	flag.StringVar(&configPath, "config", "./config/config.toml", "Use config file")
	flag.Parse()

	c := new (Config)

	if file, err := os.Open(configPath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}