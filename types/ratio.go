package t

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

/*
* 1點為0.00000001
* 在資料庫使用此欄位時，須使用BIGINT
* 在資料庫不使用NUMEIC是因為這樣容量需求比較小(8 bytes)
* MaxUint64: 9223372036854775807
 */
type Ratio uint64

func (r Ratio) Decimal() decimal.Decimal {
	d := decimal.NewFromUint64(uint64(r)).Shift(-8)
	return d
}

func (r Ratio) MultiplyAmount(amount Amount) (Amount, error) {
	result := r.Decimal().Mul(amount.Decimal)
	return NewAmountFromDecimal(result)
}
func (r Ratio) Add(arr ...Ratio) (Ratio, error) {
	d := r.Decimal()
	for _, v := range arr {
		d = d.Add(v.Decimal())
	}
	return NewRatioFromDecimal(d)
}
func (r Ratio) Sub(arr ...Ratio) (Ratio, error) {
	d := r.Decimal()
	for _, v := range arr {
		d = d.Sub(v.Decimal())
	}
	return NewRatioFromDecimal(d)
}
func (r Ratio) Mul(arr ...Ratio) (Ratio, error) {
	d := r.Decimal()
	for _, v := range arr {
		d = d.Mul(v.Decimal())
	}
	return NewRatioFromDecimal(d)
}
func MultiplyRatios(arr ...Ratio) (Ratio, error) {
	if len(arr) == 0 {
		return 0, errors.New("multiply empty Ratios")
	}
	r := arr[0]
	if len(arr) == 1 {
		return r, nil
	}
	return r.Mul(arr[1:]...)
}
func (r Ratio) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.Decimal().String()))
}

func (r *Ratio) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	d, err := decimal.NewFromString(str)
	if err != nil {
		return err
	}
	if d.IsNegative() {
		return fmt.Errorf("invalid ratio value: `%s`", d.String())
	}

	// user輸入的0.00000001 代表 1點的Ratio
	*r = Ratio(d.Shift(8).IntPart())
	return nil
}

func (r *Ratio) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Decimal().String())
}
func (r *Ratio) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	ratio, err := NewRatioFromString(str)
	if err != nil {
		return err
	}
	*r = ratio
	return nil
}

func (r Ratio) EncodeValues(key string, v *url.Values) error {
	v.Add(key, r.Decimal().String())
	return nil
}

// str: 輸入最小單位0.00000001，小於0.00000001的值將會被忽略
func NewRatioFromString(str string) (Ratio, error) {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return 0, err
	}

	return NewRatioFromDecimal(d)
}

// 最小單位0.00000001，小於0.00000001的值將會被忽略
// 最終都要由NewRatioFromDecimal來生成Ratio, 才會檢查該值的合理性
func NewRatioFromDecimal(d decimal.Decimal) (Ratio, error) {
	d = d.Shift(8)
	if !d.BigInt().IsUint64() {
		return 0, fmt.Errorf("decimal %v can not be presented as uint64", d)
	}
	i := d.IntPart()

	if d.IsNegative() {
		return 0, fmt.Errorf("invalid negative Ratio: %v", d)
	}
	return Ratio(i), nil
}
