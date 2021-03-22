package src

import (
	"gopkg.in/yaml.v3"
	"io"
)

type Config struct {
	Packages []*pkg `json:"packages" yaml:"packages"`
}

func NewConfig(r io.Reader) (*Config, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	var c Config
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
