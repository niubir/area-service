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
	Mode          string      `yaml:"mode"`
	Debug         bool        `yaml:"debug"`
	AmapKey       string      `yaml:"amapKey"`
	AutoFreshTime string      `yaml:"autoFreshTime"`
	HTTP          *httpConfig `yaml:"http"`
	GRPC          *grpcConfig `yaml:"grpc"`
}

type httpConfig struct {
	Enable bool `yaml:"enable"`
	Port   int  `yaml:"port"`
}

type grpcConfig struct {
	Enable bool `yaml:"enable"`
	Port   int  `yaml:"port"`
}

func (c *config) init() error {
	c.Mode = strings.ToUpper(c.Mode)
	if c.Mode != mode_production && c.Mode != mode_development && c.Mode != mode_test {
		c.Mode = mode_production
	}

	if c.AmapKey == "" {
		return errors.New("invalid config: amapKey")
	}

	if c.HTTP == nil {
		c.HTTP = &httpConfig{Enable: true}
	}
	if c.GRPC == nil {
		c.GRPC = &grpcConfig{Enable: true}
	}

	if err := c.HTTP.init(); err != nil {
		return err
	}

	if err := c.GRPC.init(); err != nil {
		return err
	}

	return nil
}

func (c *httpConfig) init() error {
	if !c.Enable {
		return nil
	}
	if c.Port <= 0 {
		c.Port = 10011
	}
	return nil
}

func (c *grpcConfig) init() error {
	if !c.Enable {
		return nil
	}
	if c.Port <= 0 {
		c.Port = 10012
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
