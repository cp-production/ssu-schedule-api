package api

import "github.com/cp-production/ssu-schedule-api/internal/app/store"

type Config struct {
    BindAddr string `toml:"bind_addr"`
    Store    *store.Config
}

func NewConfig() *Config {
    return &Config{
        BindAddr: ":8080",
        Store:    store.NewConfig(),
    }
}
