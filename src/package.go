package src

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

var Registry = map[string]Fetcher{}

var ErrorTypeNotSet = errors.New("type not set on package")
var ErrorTypeNotInRegistry = errors.New("type not in registry")

type Fetcher interface {
	Fetch(Package) (string, error)
}

type Package struct {
	Type     string `json:"type" yaml:"type"`
	Name     string `json:"name" yaml:"name"`
	Source   string `json:"source" yaml:"source"`
	Version  string `json:"version" yaml:"version"`
	Location string `json:"location" yaml:"location"`
}

func (p Package) WithFields() *log.Entry {
	return log.WithFields(log.Fields{
		"type": p.Type,
		"name": p.Name,
		"source": p.Source,
		"version": p.Version,
		"location": p.Location,
	})
}

type pkg struct {
	Package `json:",inline" yaml:",inline"`
	Fetcher `json:"-" yaml:"-"`
}

func copyPkg(p pkg) pkg {
	return pkg{
		Package: Package{
			Type:     p.Type,
			Name:     p.Name,
			Source:   p.Source,
			Version:  p.Version,
			Location: p.Location,
		},
		Fetcher: p.Fetcher,
	}
}

func (p *pkg) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type P pkg
	out := P{}
	err := unmarshal(&out)

	if out.Type == "" {
		return ErrorTypeNotSet
	}

	for k, v := range Registry {
		if out.Type == k {
			c := copyPkg(pkg(out))
			c.Fetcher = v
			*p = c
			return err
		}
	}
	return ErrorTypeNotInRegistry
}
