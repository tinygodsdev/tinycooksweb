package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env string `envconfig:"ENV" default:"dokku"`

	HTTPHost string `envconfig:"HTTP_HOST" default:"localhost"`
	HTTPPort string `envconfig:"HTTP_PORT" default:"8080"`

	Version string `envconfig:"VERSION" default:"v0.0.1"`

	Secret          string `envconfig:"SECRET" required:"true"`
	Expire          int    `envconfig:"EXPIRE" default:"2592000"` // seconds
	LiveSessionName string `envconfig:"LIVE_SESSION_NAME" default:"tcw-go-live-session"`

	GoogleAnalyticsID string `envconfig:"GOOGLE_ANALYTICS_ID" default:""`

	StorageType   string `envconfig:"STORAGE_TYPE" default:"ent"`
	StorageDriver string `envconfig:"STORAGE_DRIVER" default:"postgres"`
	StorageDSN    string `envconfig:"STORAGE_DSN" required:"true"`
	LogDBQueries  bool   `envconfig:"LOG_DB_QUERIES" default:"false"`
	UseMocks      bool   `envconfig:"USE_MOCKS" default:"false"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Dev() bool {
	return c.Env == "dev"
}

func (c *Config) GetAddress() string {
	return c.HTTPHost + ":" + c.HTTPPort
}
