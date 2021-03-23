package src

import (
	"gopkg.in/yaml.v3"
	"testing"
)

type customFetcher struct {
}

type customFetcher2 struct {
}

func (c customFetcher2) Fetch(p Package) (string, error) {
	return "2", nil
}

func (c customFetcher) Fetch(p Package) (string, error) {
	return "OK", nil
}

func TestPackage_UnmarshalYAML(t *testing.T) {
	Registry["custom"] = customFetcher{}
	Registry["custom2"] = customFetcher2{}
	var p Pkg
	b := []byte("name: test\ntype: custom")
	err := yaml.Unmarshal(b, &p)
	if err != nil {
		t.Error(err)
	}
	if p.Name != "test" {
		t.Error(p.Name)
	}
	if p.Type != "custom" {
		t.Error(p.Type)
	}
	s, err := p.Fetch(p.Package)
	if err != nil {
		t.Error(err)
	}
	if s != "OK" {
		t.Error(s)
	}
}
