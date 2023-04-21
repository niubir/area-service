package config

import (
	"errors"
	"strings"
)

const (
	mode_production  = "PRODUCTION"
	mode_test        = "TEST"
	mode_development = "DEVELOPMENT"
)

type config struct {
	Mode          string `yaml:"mode"`
	Debug         bool   `yaml:"debug"`
	AmapKey       string `yaml:"amapKey"`
	AutoFreshTime string `yaml:"autoFreshTime"`
}

func (c *config) init() error {
	c.Mode = strings.ToUpper(c.Mode)
	if c.Mode != mode_production && c.Mode != mode_development && c.Mode != mode_test {
		c.Mode = mode_production
	}

	if c.AmapKey == "" {
		return errors.New("invalid config: amapKey")
	}

	return nil
}

func (c config) IsProduction() bool {
	return c.Mode == mode_production
}

func (c config) IsTest() bool {
	return c.Mode == mode_test
}

func (c config) IsDevelopment() bool {
	return c.Mode == mode_development
}
