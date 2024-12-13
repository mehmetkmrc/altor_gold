package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	configPath = "config.yml"
)

type (
	Container struct {
		App      `yaml:"app"`
		Log      `yaml:"log"`
		HTTP     *HTTP     `yaml:"http"`
		PSQL     *PSQL     `yaml:"psql"`
		RabbitMQ *RabbitMQ `yaml:"rabbitmq"`
		Token    *Token    `yaml:"token"`
		Settings *Settings `yaml:"settings"`
	}

	App struct {
		Name string `env-required:"true" yaml:"name" env:"APP_NAME"`
	}

	Log struct {
		Level int `env-required:"true" yaml:"level" env:"LOG_LEVEL"`
	}

	HTTP struct {
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
		Port int    `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	Token struct {
		TokenTTL time.Duration `env-required:"true" yaml:"token_ttl" env:"TOKEN_TTL"`
		RefreshTTL time.Duration `env-required:"true" yaml:"refresh_ttl" env:"REFRESH_TTL"`
	}

	PSQL struct{
		URL string `env-required:"true" yaml:"url" env:"PSQL_URL"`
	}

	RabbitMQ struct {
		URL string `env-required:"true" yaml:"url" env:"RABBITMQ_URL"`
	}

	Settings struct{
		ServerReadTimeout int    `env-required:"true" yaml:"server_read_timeout" env:"SERVER_READ_TIMEOUT"`
		PSQLConnAttempts int    `env-required:"true" yaml:"psql_conn_attempts" env:"PSQL_CONN_ATTEMPTS"`
		PSQLConnTimeout  int    `env-required:"true" yaml:"psql_conn_timeout" env:"PSQL_CONN_TIMEOUT"`
	}
)

func NewConfig() (*Container, error) {
	cfg := new(Container)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil{
		return nil, fmt.Errorf("config error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}