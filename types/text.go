package t

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type TextArray []string

func (t *TextArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	*t = strs
	return nil
}

func (t TextArray) Value() (driver.Value, error) {
	if len(t) == 0 {
		return "{}", nil
	}
	output := strings.Join(t, ",")
	return fmt.Sprintf("{%s}", output), nil
}

// 回傳不重複的Text
func (arr *TextArray) Unique() TextArray {
	var output TextArray
	mp := map[string]bool{}
	for _, str := range *arr {
		mp[str] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}
