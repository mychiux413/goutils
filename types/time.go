package t

import (
	"database/sql/driver"
	"strings"
	"time"
)

type TimeArray []*time.Time

const TIMEZONE_LAYOUT = "2006-01-02 15:04:05-07"

func (ts *TimeArray) Scan(value interface{}) error {

	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		str = strings.Replace(str, `"`, "", 2)
		t, err := time.Parse(TIMEZONE_LAYOUT, str)
		if err != nil {
			return err
		}
		*ts = append(*ts, &t)
	}
	return nil
}

func (ts TimeArray) Value() (driver.Value, error) {
	if len(ts) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, t := range ts {

		strs = append(strs, t.Format(TIMEZONE_LAYOUT))
	}
	arr := "{" + strings.Join(strs, ",") + "}"
	return arr, nil
}
