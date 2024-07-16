package config

import (
	"time"

	"github.com/imantung/boilerplate-go-backend/internal/app/infra/di"
	"github.com/kelseyhightower/envconfig"
)

type (
	Config struct {
		Address  string         `envconfig:"ADDRESS" required:"true" default:":1323"`
		Postgres DatabaseConfig `envconfig:"PG"`
		Redis    RedisConfig    `envconfig:"REDIS"`

		BasicAuth struct {
			Username string `envconfig:"USERNAME" default:"joe" required:"true"`
			Secret   string `envconfig:"SECRET" default:"secret" required:"true"`
		} `envconfig:"BASIC_AUTH"`

		Debug bool `envconfig:"DEBUG" default:"true"`
	}

	DatabaseConfig struct {
		Source          string        `envconfig:"SOURCE" required:"true" default:"postgres://my_pg_user:my_pg_pass@localhost:5432/my_pg_dbname?sslmode=disable"`
		MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"30" required:"true"`
		MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"6" required:"true"`
		ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"30m" required:"true"`
	}

	RedisConfig struct {
		Address  string `envconfig:"ADDRESS" required:"true" default:"localhost:6379"`
		Password string `envconfig:"PASSWORD" required:"true" default:"my_redis_pass"`
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
