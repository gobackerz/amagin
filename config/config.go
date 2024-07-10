package config

import (
	"fmt"
	"github.com/gobackerz/amagin/config/errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"

	pkgLog "github.com/gobackerz/amagin/log"
)

type config struct {
	logger pkgLog.Logger
}

func New(logger pkgLog.Logger, configDir ...string) Config {
	c := &config{logger: logger}
	depEnv := fmt.Sprintf(".%v.env", c.Get("DEP_ENV", "local"))
	dir := "./configs/"

	if len(configDir) != 0 {
		dir = configDir[0]
	}

	configFile := filepath.Join(dir, depEnv)

	if depEnv != ".local.env" {
		if err := godotenv.Load(configFile); err != nil {
			c.logger.Fatalf("Failed to read %v configs: %v", depEnv, err)
		}
	}

	if err := c.readOptionalConfig(filepath.Join(dir, depEnv)); err != nil {
		c.logger.Fatalf("Failed to read .local.env configs: %v", err)
	}

	if err := c.readOptionalConfig(filepath.Join(dir, ".env")); err != nil {
		c.logger.Fatalf("Failed to read .env configs: %v", err)
	}

	return c
}

func (c *config) Get(key string, defaultVal ...string) string {
	val := os.Getenv(key)

	if strings.TrimSpace(val) == "" {
		val = c.getDefaultValue(defaultVal...)
	}

	return val
}

func (c *config) Set(key string, value string) error {
	if err := os.Setenv(key, value); err != nil {
		return errors.Config{Operation: errors.CONFIG_SET, Key: key, Err: err}
	}

	return nil
}

func (c *config) Unset(key string) error {
	if err := os.Unsetenv(key); err != nil {
		return errors.Config{Operation: errors.CONFIG_UNSET, Key: key, Err: err}
	}

	return nil
}

func (c *config) getDefaultValue(value ...string) string {
	var defaultVal string

	if len(value) > 0 {
		defaultVal = value[0]
	}

	return defaultVal
}

func (c *config) readOptionalConfig(configFile string) error {
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
