package config

import (
	"os"
	"path/filepath"

	"go.uber.org/config"
	"go.uber.org/fx"
)

var Module = fx.Provide(New)

type RuntimeEnvironment string

const (
	_envKeyPrefix             = "REACTOR_"
	_envKeyConfigDir          = "CONFIG_DIR"
	_envKeyRuntimeEnvironment = "RUNTIME_ENVIRONMENT"

	RuntimeEnvironmentDevelopment RuntimeEnvironment = "development"
	RuntimeEnvironmentProduction                     = "production"
)

type Config struct {
	RuntimeEnvironment RuntimeEnvironment `yaml:"-"`

	Controller struct {
		WorkQueueEventRetries int `yaml:"work_queue_event_retries"`
	} `yaml:"controller"`

	Kubernetes struct {
		ApiServer struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
		} `yaml:"api_server"`
	} `yaml:"kubernetes"`
}

func New() (Config, error) {
	var (
		c                  Config
		runtimeEnvironment RuntimeEnvironment
	)

	configDir := os.Getenv(_envKeyPrefix + _envKeyConfigDir)
	if configDir == "" {
		configDir = "config"
	}

	opts := []config.YAMLOption{
		config.File(filepath.Join(configDir, "base.yaml")),
	}

	switch os.Getenv(_envKeyPrefix + _envKeyRuntimeEnvironment) {
	case RuntimeEnvironmentProduction:
		runtimeEnvironment = RuntimeEnvironmentProduction
	default:
		runtimeEnvironment = RuntimeEnvironmentDevelopment
	}

	switch runtimeEnvironment {
	case RuntimeEnvironmentProduction:
		opts = append(
			opts, config.File(filepath.Join(configDir, "production.yaml")),
		)
	default:
		opts = append(
			opts, config.File(filepath.Join(configDir, "development.yaml")),
		)
	}

	opts = append(opts, config.File(filepath.Join(configDir, "secrets.yaml")))

	provider, err := config.NewYAML(opts...)
	if err != nil {
		return Config{}, err
	}

	if err := provider.Get(config.Root).Populate(&c); err != nil {
		return Config{}, err
	}

	c.RuntimeEnvironment = runtimeEnvironment

	return c, err
}
