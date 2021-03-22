package v1wporg

import (
	"errors"
	"fmt"
	"github.com/bukowa/pkgo/src"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Kind string

const KindTheme Kind = "theme"
const KindPlugin Kind = "plugin"

type Fetcher struct {
	Kind Kind
}

func (t Fetcher) Fetch(p src.Package) (string, error) {
	uri := buildWpOrgURL(wpOrgUri{
		Name: p.Source,
		Version: p.Version,
		Kind: t.Kind,
	})

	p.WithFields().Infof("fetching wporg %s %s", t.Kind, uri)
	err := urlExists(uri)
	if err != nil {
		return uri, err
	}

	resp, err := httpRequest("GET", uri, nil)
	if err != nil {
		return uri, err
	}

	defer resp.Body.Close()
	path, err := filepath.Abs(p.Destination)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL, 0660)
	if err != nil {
		return uri, err
	}

	defer f.Close()
	bWritten, err := io.Copy(f, resp.Body)
	if err != nil {
		return uri, err
	}
	if bWritten == 0 {
		return uri, errors.New("no bytes written")
	}
	return path, nil
}

type wpOrgUri struct {
	Name    string
	Version string
	Kind    Kind
}

func buildWpOrgURL(o wpOrgUri) (uri string) {
	uri = fmt.Sprintf("https://downloads.wordpress.org/%s", o.Kind)
	if o.Version == "" {
		return fmt.Sprintf("%s/%s.zip", uri, o.Name)
	}
	return fmt.Sprintf("%s/%s.%s.zip", uri, o.Name, o.Version)
}


func httpRequest(method, uri string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}

func urlExists(uri string) error {
	resp, err := httpRequest("HEAD", uri, nil)
	if err != nil {
		return err
	}
	if !isAllowedExist(resp.StatusCode) {
		return errors.New(resp.Status)
	}
	return nil
}

var allowedStatusExist = []int{
	http.StatusOK,
}
func isAllowedExist(r int) bool {
	for _, s := range allowedStatusExist {
		if s == r {
			return true
		}
	}
	return false
}