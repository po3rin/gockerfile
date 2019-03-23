package config

import (
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// Config is yml top level config.
type Config struct {
	APIVersion string `yaml:"apiVersion"`
	Repo       string `yaml:"repo"`
	Path       string `yaml:"path"`
}

// NewConfigFromBytes returns a new config from a bytes
func NewConfigFromBytes(b []byte) (*Config, error) {
	c := &Config{}
	if err := yaml.Unmarshal(b, c); err != nil {
		return nil, errors.Wrap(err, "unmarshal config")
	}
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
