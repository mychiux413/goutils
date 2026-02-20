package t

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFlow_MarshalJSON(t *testing.T) {
	d, _ := decimal.NewFromString("123.456789")
	a, _ := NewFlowFromDecimal(d)

	b, err := json.Marshal(&a)
	assert.NoError(t, err)

	// JSON 字串本身應該包含引號
	expected := `"123.456789"`
	assert.Equal(t, expected, string(b))
}

func TestFlow_UnmarshalJSON(t *testing.T) {
	jsonStr := `"987.654321"`
	var a Flow
	err := json.Unmarshal([]byte(jsonStr), &a)
	assert.NoError(t, err)

	expected, _ := decimal.NewFromString("987.654321")
	assert.True(t, a.Decimal.Equal(expected))
}

func TestFlow_MarshalGQL(t *testing.T) {
	d, _ := decimal.NewFromString("123.456789")
	a, _ := NewFlowFromDecimal(d)

	var b bytes.Buffer
	a.MarshalGQL(&b)

	// GQL 輸出應該是帶引號的字串
	expected := `"123.456789"`
	assert.Equal(t, expected, b.String())
}

func TestFlow_UnmarshalGQL(t *testing.T) {
	val := "456.789"
	var a Flow
	err := a.UnmarshalGQL(val)
	assert.NoError(t, err)

	expected, _ := decimal.NewFromString("456.789")
	assert.True(t, a.Decimal.Equal(expected))
}

func TestFlow_Precision(t *testing.T) {
	// 測試高精度數值不會遺失
	str := "1.000001215"
	d, _ := decimal.NewFromString(str)
	a, _ := NewFlowFromDecimal(d)

	// Marshal JSON
	b, err := json.Marshal(&a)
	assert.NoError(t, err)
	assert.Equal(t, `"`+str+`"`, string(b))

	// Marshal GQL
	var buf bytes.Buffer
	a.MarshalGQL(&buf)
	assert.Equal(t, `"`+str+`"`, buf.String())

	// Unmarshal JSON
	var a2 Flow
	err = json.Unmarshal([]byte(`"`+str+`"`), &a2)
	assert.NoError(t, err)
	assert.Equal(t, str, a2.String())
}
