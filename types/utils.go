package t

import (
	"fmt"
	"reflect"
	"strings"
)

func sqlStrToStrings(value interface{}) ([]string, error) {
	str, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value should be []string, but got %s", reflect.TypeOf(value))
	}
	str = str[1:(len(str) - 1)]
	if str == "" {
		return []string{}, nil
	}
	return strings.Split(str, ","), nil
}
