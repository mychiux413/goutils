package t

import (
	"database/sql/driver"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

type BigID uint64

func (id BigID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *BigID) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalID(v)
	if err != nil {
		return err
	}

	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return err
	}
	*id = BigID(i)

	return nil
}

type BigIDArray []BigID

func (ids *BigIDArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*ids = append(*ids, BigID(i))
	}
	return nil
}

func (ids BigIDArray) Value() (driver.Value, error) {
	if len(ids) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, id := range ids {
		strs = append(strs, fmt.Sprintf("%d", id))
	}
	output := strings.Join(strs, ",")
	return fmt.Sprintf("{%s}", output), nil
}

func (arr BigIDArray) MarshalGQL(w io.Writer) {
	var strs []string
	for _, id := range arr {
		strs = append(strs, strconv.Quote(fmt.Sprintf("%d", id)))
	}
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(strs, ",")))
}

func (id *BigIDArray) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("BigIDArray must be a string")
	}
	if len(value) < 2 {
		return fmt.Errorf("invalid array: %s", value)
	}

	value = value[1 : len(value)-1]
	for _, str := range strings.Split(value, ",") {
		str, err := strconv.Unquote(str)
		if err != nil {
			return err
		}
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*id = append(*id, BigID(i))
	}

	return nil
}

func (ids *BigIDArray) Includes(target BigID) bool {
	for _, id := range *ids {
		if id == target {
			return true
		}
	}
	return false
}

type ID uint32

func (id ID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *ID) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalID(v)
	if err != nil {
		return err
	}

	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return err
	}
	*id = ID(i)

	return nil
}

type IDArray []ID

func (ids *IDArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*ids = append(*ids, ID(i))
	}
	return nil
}

func (ids IDArray) Value() (driver.Value, error) {
	if len(ids) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, id := range ids {
		strs = append(strs, fmt.Sprintf("%d", id))
	}
	output := strings.Join(strs, ",")
	return fmt.Sprintf("{%s}", output), nil
}

func (arr IDArray) MarshalGQL(w io.Writer) {
	var strs []string
	for _, id := range arr {
		strs = append(strs, strconv.Quote(fmt.Sprintf("%d", id)))
	}
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(strs, ",")))
}

func (id *IDArray) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("IDArray must be a string")
	}
	if len(value) < 2 {
		return fmt.Errorf("invalid array: %s", value)
	}

	value = value[1 : len(value)-1]
	for _, str := range strings.Split(value, ",") {
		str, err := strconv.Unquote(str)
		if err != nil {
			return err
		}
		i, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return err
		}
		*id = append(*id, ID(i))
	}

	return nil
}

func (ids *IDArray) Includes(target ID) bool {
	for _, id := range *ids {
		if id == target {
			return true
		}
	}
	return false
}

type SmallID uint16

func (id SmallID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *SmallID) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalID(v)
	if err != nil {
		return err
	}

	i, err := strconv.ParseInt(str, 10, 16)
	if err != nil {
		return err
	}
	*id = SmallID(i)

	return nil
}

type SmallIDArray []SmallID

func (ids *SmallIDArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*ids = append(*ids, SmallID(i))
	}
	return nil
}

func (ids SmallIDArray) Value() (driver.Value, error) {
	if len(ids) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, id := range ids {
		strs = append(strs, fmt.Sprintf("%d", id))
	}
	output := strings.Join(strs, ",")
	return fmt.Sprintf("{%s}", output), nil
}

func (arr SmallIDArray) MarshalGQL(w io.Writer) {
	var strs []string
	for _, id := range arr {
		strs = append(strs, strconv.Quote(fmt.Sprintf("%d", id)))
	}
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(strs, ",")))
}

func (id *SmallIDArray) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("IDArray must be a string")
	}
	if len(value) < 2 {
		return fmt.Errorf("invalid array: %s", value)
	}

	value = value[1 : len(value)-1]
	for _, str := range strings.Split(value, ",") {
		str, err := strconv.Unquote(str)
		if err != nil {
			return err
		}
		i, err := strconv.ParseUint(str, 10, 16)
		if err != nil {
			return err
		}
		*id = append(*id, SmallID(i))
	}

	return nil
}

func (ids *SmallIDArray) Includes(target SmallID) bool {
	for _, id := range *ids {
		if id == target {
			return true
		}
	}
	return false
}

type TinyID uint8

func (id TinyID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *TinyID) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalID(v)
	if err != nil {
		return err
	}

	i, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return err
	}
	*id = TinyID(i)

	return nil
}

type TinyIDArray []TinyID

func (ids *TinyIDArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		i, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			return err
		}
		*ids = append(*ids, TinyID(i))
	}
	return nil
}

func (ids TinyIDArray) Value() (driver.Value, error) {
	if len(ids) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, id := range ids {
		strs = append(strs, fmt.Sprintf("%d", id))
	}
	output := strings.Join(strs, ",")
	return fmt.Sprintf("{%s}", output), nil
}

func (arr TinyIDArray) MarshalGQL(w io.Writer) {
	var strs []string
	for _, id := range arr {
		strs = append(strs, strconv.Quote(fmt.Sprintf("%d", id)))
	}
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(strs, ",")))
}

func (id *TinyIDArray) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("IDArray must be a string")
	}
	if len(value) < 2 {
		return fmt.Errorf("invalid array: %s", value)
	}

	value = value[1 : len(value)-1]
	for _, str := range strings.Split(value, ",") {
		str, err := strconv.Unquote(str)
		if err != nil {
			return err
		}
		i, err := strconv.ParseUint(str, 10, 32)
		if err != nil {
			return err
		}
		*id = append(*id, TinyID(i))
	}

	return nil
}

func (ids *TinyIDArray) Includes(target TinyID) bool {
	for _, id := range *ids {
		if id == target {
			return true
		}
	}
	return false
}
