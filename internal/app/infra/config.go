package infra

import (
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		ServerAddress string `envconfig:"SERVER_ADDRESS" default:":1323"`
		DatabaseURL   string `envconfig:"DATABASE_URL" default:"postgres://postgres:postgres@localhost:5432/database?sslmode=disable"`

		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"30" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"6" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"30m" required:"true"`
	}
)

var _ = di.Provide(NewConfig)

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
