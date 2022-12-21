package main

import (
	"fmt"

	oosConfig "wemade_project/config"
)

//main
func main() {
	config := oosConfig.GetConfig()

	// var port int
	// var configPath string
	// flag.IntVar(&port, "port", 7080, "port to listen on")

	// flag.StringVar(&configPath, "config", "./config/config1.toml", "Use config file")

	// flag.Parse()

	// fmt.Print("Hello World Go Go = ", port )
	fmt.Print("Hello World Go Go = ", config.DB.Host )
	// fmt.Println("configPath = ", configPath)
	// flag.Parse()
	// fmt.Println(flag.Args())
}