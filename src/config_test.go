package src

import (
	"bytes"
	"testing"
)

func TestNewConfig(t *testing.T) {
	Registry["custom"] = customFetcher{}
	Registry["custom2"] = customFetcher2{}

	r := bytes.NewBuffer([]byte(`packages:
  - name: test
    version: 1
    type: custom
  - name: test2
    version: 2
    type: custom2`))

	c, err := NewConfig(r)
	if err != nil {
		t.Error(err)
	}
	p1, p2 := c.Packages[0], c.Packages[1]
	if p1.Type != "custom" || p2.Type != "custom2"{
		t.Error()
	}
	if p1.Version != "1" || p2.Version != "2" {
		t.Error()
	}

	if s, err := p1.Fetch(p1.Package); s != "OK" || err != nil {
		t.Error(err)
	}

	if s, err := p2.Fetch(p1.Package); s != "2" || err != nil {
		t.Error(err)
	}

}
