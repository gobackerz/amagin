package amagin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	"github.com/gobackerz/amagin"
	"github.com/gobackerz/amagin/errors"
)

type Config struct {
	logger amagin.Logger
}

func newConfig(logger amagin.Logger, configDir ...string) *Config {
	c := &Config{logger: logger}
	depEnv := fmt.Sprintf(".%v.env", c.Get("DEP_ENV", "local"))
	dir := "./configs/"

	if len(configDir) != 0 {
		dir = configDir[0]
	}

	configFile := filepath.Join(dir, depEnv)

	if depEnv != ".local.env" {
		if err := godotenv.Load(configFile); err != nil {
			c.logger.Error("Failed to read %v configs: %v", depEnv, err)
			os.Exit(0)
		}
	}

	if err := c.readOptionalConfig(filepath.Join(dir, depEnv)); err != nil {
		c.logger.Error("Failed to read .local.env configs: %v", err)
		os.Exit(0)
	}

	if err := c.readOptionalConfig(filepath.Join(dir, ".env")); err != nil {
		c.logger.Error("Failed to read .env configs: %v", err)
		os.Exit(0)
	}

	return c
}

func (c *Config) Get(key string, defaultVal ...string) string {
	val := os.Getenv(key)

	if strings.TrimSpace(val) == "" {
		val = c.getDefaultValue(defaultVal...)
	}

	return val
}

func (c *Config) Set(key string, value string) error {
	if err := os.Setenv(key, value); err != nil {
		return errors.Config{Operation: errors.CONFIG_SET, Key: key, Err: err}
	}

	return nil
}

func (c *Config) Unset(key string) error {
	if err := os.Unsetenv(key); err != nil {
		return errors.Config{Operation: errors.CONFIG_UNSET, Key: key, Err: err}
	}

	return nil
}

func (c *Config) getDefaultValue(value ...string) string {
	var defaultVal string

	if len(value) > 0 {
		defaultVal = value[0]
	}

	return defaultVal
}

func (c *Config) readOptionalConfig(configFile string) error {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}

	if err = godotenv.Load(configFile); err != nil {
		return err
	}

	return nil
}
