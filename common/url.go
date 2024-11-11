package c

import (
	"fmt"
	"net/url"
	"path"
	"sort"
	"strings"

	"github.com/google/go-querystring/query"
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

func IsURLQueryString(str string) bool {
	return str == url.QueryEscape(str)
}

// keyDesc: true => 則所有的key都會降冪排列，反之升冪
// 使用的tag為 `url`
func ToQueryString(obj interface{}, keyDesc bool) (string, error) {
	values, err := query.Values(obj)
	if err != nil {
		return "", err
	}

	var keys []string
	for k := range values {
		keys = append(keys, k)
	}

	if keyDesc {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	} else {
		sort.Strings(keys)
	}
	var strs []string
	for _, key := range keys {
		value := values.Get(key)
		strs = append(strs, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
	}

	return strings.Join(strs, "&"), nil
}
