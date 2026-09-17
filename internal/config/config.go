package config

import (
	"time"
)

type Config struct {
	RefreshInterval time.Duration
	Theme           string
	ProcessLimit    int
}

func Default() Config {
	return Config{
		RefreshInterval: time.Second,
		Theme:           "dark",
		ProcessLimit:    50,
	}
}
