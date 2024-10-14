package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

/*
 * 在資料庫使用此欄位時，要使用NUMRIC避免精度消失
 * 只要金錢相關的欄位都應該用此type，避免計算過程有精度消失的問題
 */
type Amount struct {
	decimal.Decimal
}

func (a *Amount) Scan(value interface{}) error {

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Amount should be string, but got %s", reflect.TypeOf(value))
	}
	amount, err := NewAmountFromString(str)
	if err != nil {
		return err
	}
	*a = *amount
	return nil
}

func (a Amount) Value() (driver.Value, error) {
	return a.String(), nil
}

func (a Amount) MarshalGQL(w io.Writer) {
	io.WriteString(w, a.String())
}

func (a Amount) Add(y Amount) Amount {
	added := a.Decimal.Add(y.Decimal)
	return Amount{
		Decimal: added,
	}
}
func (a Amount) Sub(y Amount) Amount {
	added := a.Decimal.Sub(y.Decimal)
	return Amount{
		Decimal: added,
	}
}
func (a *Amount) Equal(y Amount) bool {
	return a.Decimal.Equal(y.Decimal)
}

func (a *Amount) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	amount, err := NewAmountFromString(str)
	if err != nil {
		return err
	}
	*a = *amount
	return nil
}

func (a *Amount) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}
func (a *Amount) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	amount, err := NewAmountFromString(str)
	if err != nil {
		return err
	}
	*a = *amount
	return nil
}

func (a Amount) EncodeValues(key string, v *url.Values) error {
	v.Add(key, a.String())
	return nil
}

func NewAmountFromDecimal(d decimal.Decimal) *Amount {
	return &Amount{
		Decimal: d,
	}
}

func NewAmountFromString(str string) (*Amount, error) {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return nil, err
	}
	return &Amount{
		Decimal: d,
	}, nil
}
