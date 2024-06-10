package config

import "github.com/amnestia/tnderlike/internal/database"

// Config struct containing config
type Config struct {
	App    string `yaml:"name"`
	Server Server `yaml:"server"`

	Environment string          `json:"environment"`
	Database    database.Config `json:"database"`
	Auth        Auth            `json:"auth"`
}

// Server struct containing server config
type Server struct {
	Port string `yaml:"port"`
	Logs struct {
		Info  string `yaml:"info"`
		Error string `yaml:"error"`
	} `yaml:"logs"`
	Timeout int64 `yaml:"timeout"`
	Limiter struct {
		Rate       int64 `yaml:"rate"`
		Burst      int   `yaml:"burst"`
		Expiration int64 `yaml:"expiration"`
	}
	Origin []string `yaml:"origin"`
}

// Auth struct containing auth key
type Auth struct {
	PubKey   string `json:"pub_key"`
	PrivKey  string `json:"priv_key"`
	LocalKey string `json:"local_key"`
	Pepper   string `json:"pepper"`
}
