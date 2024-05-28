package config

import (
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Address  string   `envconfig:"ADDRESS" required:"true" default:":1323"`
		Postgres DBConfig `envconfig:"PG"`
	}
	DBConfig struct {
		DBName string `envconfig:"DBNAME" required:"true" default:"postgres"`
		DBUser string `envconfig:"DBUSER" required:"true" default:"postgres"`
		DBPass string `envconfig:"DBPASS" required:"true" default:"postgres"`
		Host   string `envconfig:"HOST" required:"true" default:"localhost"`
		Port   string `envconfig:"PORT" required:"true" default:"5432"`

		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"30" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"6" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"30m" required:"true"`
	}
)

const Prefix = "APP"

var _ = di.Provide(NewConfig)

func NewConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process(Prefix, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
