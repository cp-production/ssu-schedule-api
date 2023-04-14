package api

type Config struct {
	BindAddr string `toml:"bind_addr"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",  
	}
}
