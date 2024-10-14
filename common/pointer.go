package c

import "reflect"

func PointerBool(value bool) *bool {
	return &value
}

func PointerString(value string) *string {
	return &value
}

func PointerUint32(value uint32) *uint32 {
	return &value
}
func PointerUint16(value uint16) *uint16 {
	return &value
}

func PointerFloat32(value float32) *float32 {
	return &value
}
func PointerUint64(value uint64) *uint64 {
	return &value
}
func PointerInt64(value int64) *int64 {
	return &value
}

func PointerFloat64(value float64) *float64 {
	return &value
}

func PointerUint(value uint) *uint {
	return &value
}

func PointerInt(value int) *int {
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
