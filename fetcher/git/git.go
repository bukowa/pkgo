package git

import "github.com/bukowa/pkgo/src"

type Fetcher struct {
}

func (f Fetcher) Fetch(p src.Package) (string, error) {
	return "git", nil
}
