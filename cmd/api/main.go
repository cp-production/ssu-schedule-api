package main

import (
	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/cp-production/ssu-schedule-api/docs"
	"github.com/cp-production/ssu-schedule-api/internal/app/api"
)

// @title SSU Schedule API
// @version 1.0
// @description API Server for SSU Schedule Application

// @host      localhost:8080
// @BasePath  /api/v1.0
func main() {
	configPath := "configs/api.toml"
	config := api.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}
	if err := api.Start(config); err != nil {
		log.Fatal(err)
	}
}
