package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

type JSONMap map[string]interface{}

func (m *JSONMap) Scan(value interface{}) error {

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("JSONMap should be bytes, but got %s", reflect.TypeOf(value))
	}
	return json.Unmarshal(bytes, m)
}

func (m JSONMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m JSONMap) MarshalGQL(w io.Writer) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Printf("Marshal JSONMap error: %v\n", err)
	}
	io.WriteString(w, strconv.Quote(string(data)))
}

func (m *JSONMap) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), m)
}

type JSONList []interface{}

func (m *JSONList) Scan(value interface{}) error {

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("JSONList should be bytes, but got %s", reflect.TypeOf(value))
	}
	return json.Unmarshal(bytes, m)
}

func (m JSONList) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m JSONList) MarshalGQL(w io.Writer) {
	data, err := json.Marshal(m)
	if err != nil {
		log.Printf("Marshal JSONList error: %v\n", err)
	}
	io.WriteString(w, strconv.Quote(string(data)))
}

func (m *JSONList) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(str), m)
}
