package config

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Config is yml top level config.
type Config struct {
	Repo string `yaml:"repo"`
	Ver  string `yaml:"version"`
	Path string `yaml:"path"`
}

func validateRequiredConfig(c *Config) error {
	if c.Repo == "" {
		return errors.New("repo field is required")
	}
	return nil
}

func SetDefaultConfig(c *Config) *Config {
	if c.Ver == "" {
		c.Ver = "master"
	}
	if c.Path == "" {
		c.Path = "./"
	}
	return c
}

// NewConfigFromBytes returns a new config from a bytes
func NewConfigFromBytes(b []byte) (*Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.Wrap(err, "unmarshal config")
	}
	if err := validateRequiredConfig(c); err != nil {
		return nil, errors.Wrap(err, "failed to validate gockerfile")
	}
	c = SetDefaultConfig(c)
	return c, nil
}

// NewConfigFromFilename shorthand to NewConfigFromBytes.
func NewConfigFromFilename(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "got error opening file")
	}
	defer f.Close()
	contents, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "got error reading config file")
	}
	return NewConfigFromBytes(contents)
}
