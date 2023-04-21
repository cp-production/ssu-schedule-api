package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/cp-production/ssu-schedule-api/internal/app/api"
)

func main() {
	configPath := "configs/api.toml"
	config := api.NewConfig()
	if _, err := toml.DecodeFile(configPath, config); err != nil {
		log.Fatal(err)
	}
    
	s := api.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
    }
}
