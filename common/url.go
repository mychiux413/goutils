package c

import (
	"net/url"
	"path"
)

func URLJoin(baseUrl string, paths ...string) (*url.URL, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	for _, p := range paths {
		u.Path = path.Join(u.Path, p)
	}
	return u, nil
}
