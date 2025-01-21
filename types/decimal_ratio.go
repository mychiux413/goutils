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
使用NUMERIC儲存的Ratio
*/
type DecimalRatio struct {
	decimal.Decimal
}

func (r DecimalRatio) MultiplyAmount(amount Amount) (Amount, error) {
	result := r.Decimal.Mul(amount.Decimal)
	return NewAmountFromDecimal(result)
}
func (r DecimalRatio) Add(arr ...DecimalRatio) (DecimalRatio, error) {
	d := r.Decimal
	for _, v := range arr {
		d = d.Add(v.Decimal)
	}
	return NewDecimalRatioFromDecimal(d)
}
func (r DecimalRatio) Sub(arr ...DecimalRatio) (DecimalRatio, error) {
	d := r.Decimal
	for _, v := range arr {
		d = d.Sub(v.Decimal)
	}
	return NewDecimalRatioFromDecimal(d)
}
func (r DecimalRatio) Mul(arr ...DecimalRatio) (DecimalRatio, error) {
	d := r.Decimal
	for _, v := range arr {
		d = d.Mul(v.Decimal)
	}
	return NewDecimalRatioFromDecimal(d)
}
func MultiplyDecimalRatios(arr ...DecimalRatio) (DecimalRatio, error) {
	if len(arr) == 0 {
		return DecimalRatio{}, errors.New("multiply empty DecimalRatios")
	}
	r := arr[0]
	if len(arr) == 1 {
		return r, nil
	}
	return r.Mul(arr[1:]...)
}
func (r DecimalRatio) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(r.Decimal.String()))
}

func (r *DecimalRatio) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	d, err := decimal.NewFromString(str)
	if err != nil {
		return err
	}
	if d.IsNegative() {
		return fmt.Errorf("invalid DecimalRatio value: `%s`", d.String())
	}

	r.Decimal = d
	return nil
}

func (r *DecimalRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.Decimal.String())
}
func (r *DecimalRatio) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	DecimalRatio, err := NewDecimalRatioFromString(str)
	if err != nil {
		return err
	}
	*r = DecimalRatio
	return nil
}

func (r DecimalRatio) EncodeValues(key string, v *url.Values) error {
	v.Add(key, r.Decimal.String())
	return nil
}

func NewDecimalRatioFromString(str string) (DecimalRatio, error) {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return DecimalRatio{}, err
	}

	return NewDecimalRatioFromDecimal(d)
}

func NewDecimalRatioFromDecimal(d decimal.Decimal) (DecimalRatio, error) {
	if d.IsNegative() {
		return DecimalRatio{}, fmt.Errorf("invalid negative DecimalRatio: %v", d)
	}
	return DecimalRatio{
		Decimal: d,
	}, nil
}
