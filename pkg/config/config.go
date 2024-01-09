package configs

import (
	"os"
	"time"

	"github.com/go-yaml/yaml"

	"nam_0801/pkg/db/postgres"
)

type Config struct {
	Postgres postgres.DatabaseConfig `yaml:"postgres"`
	Server   Server                  `yaml:"server"`
}

type Server struct {
	Address         string        `yaml:"address"`
	ReadTimeout     time.Duration `yaml:"read_timeout"`
	WriteTimeout    time.Duration `yaml:"write_timeout"`
	ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
	CORS            struct {
		AllowOrigins []string `yaml:"allowOrigins"`
		AllowMethods []string `yaml:"allowMethods"`
		AllowHeaders []string `yaml:"allowHeaders"`
	} `yaml:"cors"`
	SkipLogPaths         []string `yaml:"skipLoggingPaths"`
	ResponseLogLimitByte int      `yaml:"responseLogLimitByte"`
	UploadSizeLimitByte  int64    `yaml:"uploadSizeLimitByte"`
	Token                Token    `yaml:"token"`
}

type Token struct {
	JwtSecretKey  string        `yaml:"jwtSecretKey"`
	TimeToExpired time.Duration `yaml:"timeToExpired"`
}

// NewConfig returns a new decoded Config struct
func NewConfig() (*Config, error) {
	// Create config structure
	config := &Config{}

	// Open config file
	file, err := os.Open("config/config.yml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	// Start YAML decoding from file
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
