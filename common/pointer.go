package c

import (
	"reflect"
	"time"
)

func PointerBool(value bool) *bool {
	return &value
}

func PointerString(value string) *string {
	return &value
}

func PointerUint64(value uint64) *uint64 {
	return &value
}
func PointerUint32(value uint32) *uint32 {
	return &value
}
func PointerUint16(value uint16) *uint16 {
	return &value
}
func PointerUint8(value uint8) *uint8 {
	return &value
}
func PointerUint(value uint) *uint {
	return &value
}

func PointerInt64(value int64) *int64 {
	return &value
}
func PointerInt32(value int32) *int32 {
	return &value
}
func PointerInt16(value int16) *int16 {
	return &value
}
func PointerInt8(value int8) *int8 {
	return &value
}
func PointerInt(value int) *int {
	return &value
}

func PointerFloat64(value float64) *float64 {
	return &value
}
func PointerFloat32(value float32) *float32 {
	return &value
}

func PointerTime(value time.Time) *time.Time {
	return &value
}

// 移除map裡的nil值
func RemoveNilValue(data map[string]interface{}) {
	for k, v := range data {
		if reflect.ValueOf(v).IsNil() {
			delete(data, k)
		}
	}
}

// ValueOrDefault 是一個泛型函式，它接受一個指標 T 和一個預設值 T。
// 如果指標為 nil，則回傳預設值。
// 如果指標不為 nil，則回傳指標所指向的值。
//
// 範例:
//
//	strPtr := PointerString("hello")
//	ValueOrDefault(strPtr, "world") // 回傳 "hello"
//
//	var nilStrPtr *string
//	ValueOrDefault(nilStrPtr, "world") // 回傳 "world"
func ValueOrDefault[T any](value *T, defaultValue T) T {
	if value == nil || reflect.ValueOf(value).IsNil() {
		return defaultValue
	}
	return *value
}
