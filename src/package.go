package src

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

// Registry is a map of Package.Type to Fetcher implementing that Type.
var Registry = map[string]Fetcher{}

var ErrorTypeNotSet = errors.New("type not set on package")
var ErrorTypeNotInRegistry = errors.New("type not in registry")

type Fetcher interface {
	Fetch(Package) (string, error)
}

type Package struct {
	Type        string `json:"type" yaml:"type"`
	Name        string `json:"name" yaml:"name"`
	Source      string `json:"source" yaml:"source"`
	Version     string `json:"version" yaml:"version"`
	Destination string `json:"destination" yaml:"destination"`
	Labels map[string]interface{}	`json:"labels" yaml:"labels"`
}

func (p Package) WithFields() *log.Entry {
	return log.WithFields(log.Fields{
		"type": p.Type,
		"name": p.Name,
		"source": p.Source,
		"version": p.Version,
		"destination": p.Destination,
		"labels": p.Labels,
	})
}

type Pkg struct {
	Package `json:",inline" yaml:",inline"`
	Fetcher `json:"-" yaml:"-"`
}

func copyPkg(p Pkg) Pkg {
	return Pkg{
		Package: Package{
			Type:        p.Type,
			Name:        p.Name,
			Source:      p.Source,
			Version:     p.Version,
			Destination: p.Destination,
			Labels: p.Labels,
		},
		Fetcher: p.Fetcher,
	}
}

func (p *Pkg) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type P Pkg
	out := P{}
	err := unmarshal(&out)

	if out.Type == "" {
		return ErrorTypeNotSet
	}

	for k, v := range Registry {
		if out.Type == k {
			c := copyPkg(Pkg(out))
			c.Fetcher = v
			*p = c
			return err
		}
	}
	return ErrorTypeNotInRegistry
}
