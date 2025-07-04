package t

import (
	"database/sql/driver"
	"fmt"
	"net/url"
	"strconv"
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

// Ex: "1575856875" (seconds)
type TimestampSec struct {
	time.Time
}

func (t TimestampSec) EncodeValues(key string, v *url.Values) error {
	j, err := t.MarshalJSON()
	if err != nil {
		return err
	}
	v.Add(key, string(j))
	return nil
}

func (dt *TimestampSec) UnmarshalJSON(bytes []byte) error {
	i, err := strconv.ParseInt(string(bytes), 10, 64)
	if err != nil {
		return err
	}

	dt.Time = time.Unix(i, 0)
	return nil
}

func (dt *TimestampSec) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("%d", dt.Unix())
	return []byte(str), nil
}

// Ex: "1575856875000" (milli seconds)
type TimestampMs struct {
	time.Time
}

func (t TimestampMs) EncodeValues(key string, v *url.Values) error {
	j, err := t.MarshalJSON()
	if err != nil {
		return err
	}
	v.Add(key, string(j))
	return nil
}

func (dt *TimestampMs) UnmarshalJSON(bytes []byte) error {
	i, err := strconv.Atoi(string(bytes))
	if err != nil {
		return err
	}

	dt.Time = time.UnixMilli(int64(i))
	return nil
}

func (dt *TimestampMs) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("%d", dt.UnixMilli())
	return []byte(str), nil
}
