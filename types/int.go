package t

import (
	"database/sql/driver"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"

	"slices"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

// 在資料庫裡，需使用NUMERIC型別
type HugeID decimal.Decimal

func NewHugeIDFromDecimal(d decimal.Decimal) (HugeID, error) {
	output := HugeID(d)
	if !output.IsValid() {
		return HugeID{}, fmt.Errorf("invalid HugeID: %s", d.String())
	}
	return HugeID(d), nil
}
func NewHugeIDFromString(str string) (HugeID, error) {
	dec, err := decimal.NewFromString(str)
	if err != nil {
		return HugeID{}, err
	}
	return NewHugeIDFromDecimal(dec)
}

// 如果有計算用途，可以轉成Deciaml
func (id HugeID) Decimal() decimal.Decimal {
	return decimal.Decimal(id)
}

func (id HugeID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(decimal.Decimal(id).String()))
}

func (a HugeID) IsValid() bool {
	dec := a.Decimal()
	if !dec.IsInteger() {
		return false
	}
	return !dec.IsNegative()
}

func (a *HugeID) Scan(value any) error {

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("HugeID should be string, but got %s", reflect.TypeOf(value))
	}
	id, err := NewHugeIDFromString(str)
	if err != nil {
		return err
	}
	*a = id
	return nil
}

func (a HugeID) Value() (driver.Value, error) {
	return a.String(), nil
}

func (id *HugeID) UnmarshalGQL(v any) error {
	str, err := graphql.UnmarshalID(v)
	if err != nil {
		return err
	}

	hugeID, err := NewHugeIDFromString(str)
	if err != nil {
		return err
	}
	*id = hugeID

	return nil
}
func (id HugeID) String() string {
	return decimal.Decimal(id).String()
}
func (id HugeID) GoString() string {
	return decimal.Decimal(id).String()
}

type HugeIDArray []HugeID

func (ids *HugeIDArray) Scan(value any) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		hugeID, err := NewHugeIDFromString(str)
		if err != nil {
			return err
		}
		*ids = append(*ids, hugeID)
	}
	return nil
}

func (ids HugeIDArray) Value() (driver.Value, error) {
	if len(ids) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, id := range ids {
		strs = append(strs, id.String())
	}
	output := strings.Join(strs, ",")
	return fmt.Sprintf("{%s}", output), nil
}

func (arr HugeIDArray) MarshalGQL(w io.Writer) {
	var strs []string
	for _, id := range arr {
		strs = append(strs, strconv.Quote(id.String()))
	}
	io.WriteString(w, fmt.Sprintf("[%s]", strings.Join(strs, ",")))
}

func (id *HugeIDArray) UnmarshalGQL(v any) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("HugeIDArray must be a string")
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
		hugeID, err := NewHugeIDFromString(str)
		if err != nil {
			return err
		}
		*id = append(*id, hugeID)
	}

	return nil
}
func (ids HugeIDArray) String() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}
func (ids HugeIDArray) GoString() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.GoString())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}

func (ids HugeIDArray) StringArray() []string {
	var strs []string
	for _, id := range ids {
		strs = append(strs, id.String())
	}
	return strs

}

func (ids *HugeIDArray) Includes(target HugeID) bool {
	arr := ids.StringArray()
	return slices.Contains(arr, target.String())
}

// 回傳不重複的ID
func (ids *HugeIDArray) Unique() HugeIDArray {
	var output HugeIDArray
	mp := map[string]HugeID{}
	for _, id := range *ids {
		mp[id.String()] = id
	}
	for _, id := range mp {
		output = append(output, id)
	}
	return output
}

type BigID uint64

func (id BigID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *BigID) UnmarshalGQL(v any) error {
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
func (id BigID) String() string {
	return fmt.Sprintf("%d", id)
}
func (id BigID) GoString() string {
	return fmt.Sprintf("%d", id)
}

type BigIDArray []BigID

func (ids *BigIDArray) Scan(value any) error {
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

func (id *BigIDArray) UnmarshalGQL(v any) error {
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
func (ids BigIDArray) String() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}
func (ids BigIDArray) GoString() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.GoString())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}

func (ids *BigIDArray) Includes(target BigID) bool {
	return slices.Contains(*ids, target)
}

// 回傳不重複的ID
func (ids *BigIDArray) Unique() BigIDArray {
	var output BigIDArray
	mp := map[BigID]bool{}
	for _, id := range *ids {
		mp[id] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}

type ID uint32

func (id ID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *ID) UnmarshalGQL(v any) error {
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

func (id ID) String() string {
	return fmt.Sprintf("%d", id)
}
func (id ID) GoString() string {
	return fmt.Sprintf("%d", id)
}

type IDArray []ID

func (ids *IDArray) Scan(value any) error {
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

func (id *IDArray) UnmarshalGQL(v any) error {
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

func (ids IDArray) String() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}
func (ids IDArray) GoString() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.GoString())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}

func (ids *IDArray) Includes(target ID) bool {
	return slices.Contains(*ids, target)
}

// 回傳不重複的ID
func (ids *IDArray) Unique() IDArray {
	var output IDArray
	mp := map[ID]bool{}
	for _, id := range *ids {
		mp[id] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}

type SmallID uint16

func (id SmallID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *SmallID) UnmarshalGQL(v any) error {
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

func (id SmallID) String() string {
	return fmt.Sprintf("%d", id)
}
func (id SmallID) GoString() string {
	return fmt.Sprintf("%d", id)
}

type SmallIDArray []SmallID

func (ids *SmallIDArray) Scan(value any) error {
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

func (id *SmallIDArray) UnmarshalGQL(v any) error {
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

func (ids SmallIDArray) String() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}
func (ids SmallIDArray) GoString() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.GoString())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}

func (ids *SmallIDArray) Includes(target SmallID) bool {
	return slices.Contains(*ids, target)
}

// 回傳不重複的ID
func (ids *SmallIDArray) Unique() SmallIDArray {
	var output SmallIDArray
	mp := map[SmallID]bool{}
	for _, id := range *ids {
		mp[id] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}

type TinyID uint8

func (id TinyID) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(strconv.Itoa(int(id))))
}

func (id *TinyID) UnmarshalGQL(v any) error {
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

func (id TinyID) String() string {
	return fmt.Sprintf("%d", id)
}
func (id TinyID) GoString() string {
	return fmt.Sprintf("%d", id)
}

type TinyIDArray []TinyID

func (ids *TinyIDArray) Scan(value any) error {
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

func (id *TinyIDArray) UnmarshalGQL(v any) error {
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

func (ids TinyIDArray) String() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.String())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}
func (ids TinyIDArray) GoString() string {
	var arr []string
	for _, id := range ids {
		arr = append(arr, id.GoString())
	}
	return fmt.Sprintf("[%s]", strings.Join(arr, ","))
}

func (ids *TinyIDArray) Includes(target TinyID) bool {
	return slices.Contains(*ids, target)
}

// 回傳不重複的ID
func (ids *TinyIDArray) Unique() TinyIDArray {
	var output TinyIDArray
	mp := map[TinyID]bool{}
	for _, id := range *ids {
		mp[id] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}

func MarshalUint8(value uint8) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%d", value))
	})
}

func UnmarshalUint8(v any) (uint8, error) {
	ui, err := graphql.UnmarshalUint(v)
	if err != nil {
		return 0, err
	}
	if ui > math.MaxUint8 {
		return 0, fmt.Errorf("invalid Uint8: %d", ui)
	}

	return uint8(ui), nil
}
