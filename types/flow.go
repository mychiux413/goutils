package t

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	c "github.com/mychiux413/goutils/common"
	"github.com/shopspring/decimal"
)

/*
 * 在資料庫使用此欄位時，要使用NUMRIC避免精度消失
 * 只要金錢相關的欄位都應該用此type，避免計算過程有精度消失的問題
 * Flow可以是負數
 */
type Flow struct {
	decimal.Decimal
}

func (a *Flow) Scan(value interface{}) error {

	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("Flow should be string, but got %s", reflect.TypeOf(value))
	}
	Flow, err := NewFlowFromString(str)
	if err != nil {
		return err
	}
	*a = Flow
	return nil
}

func (a Flow) Value() (driver.Value, error) {
	return a.String(), nil
}

func (a Flow) MarshalGQL(w io.Writer) {
	graphql.MarshalString(a.String()).MarshalGQL(w)
}

func (a Flow) Add(arr ...Flow) (Flow, error) {
	d := a.Decimal
	for _, Flow := range arr {
		d = d.Add(Flow.Decimal)
	}
	return NewFlowFromDecimal(d)
}
func (a Flow) Sub(arr ...Flow) (Flow, error) {
	d := a.Decimal
	for _, Flow := range arr {
		d = d.Sub(Flow.Decimal)
	}
	return NewFlowFromDecimal(d)
}

// 如果在算式裡執行了這個，肯定是哪裡寫錯了！！
func (a Flow) Mul(y Flow) (Flow, error) {
	return Flow{}, fmt.Errorf("%w: dont multiply Flow by Flow: %v x %v", c.ErrBadRequest, a, y)
}
func (a *Flow) Equal(y Flow) bool {
	return a.Decimal.Equal(y.Decimal)
}
func (a *Flow) UnmarshalGQL(v interface{}) error {
	str, err := graphql.UnmarshalString(v)
	if err != nil {
		return err
	}

	Flow, err := NewFlowFromString(str)
	if err != nil {
		return err
	}
	*a = Flow
	return nil
}

func (a *Flow) MarshalJSON() ([]byte, error) {
	return json.Marshal(a.String())
}
func (a *Flow) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	Flow, err := NewFlowFromString(str)
	if err != nil {
		return err
	}
	*a = Flow
	return nil
}

func (a Flow) EncodeValues(key string, v *url.Values) error {
	v.Add(key, a.String())
	return nil
}

func NewFlowFromDecimal(d decimal.Decimal) (Flow, error) {
	return Flow{
		Decimal: d,
	}, nil
}

func NewFlowFromString(str string) (Flow, error) {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return Flow{}, err
	}
	return NewFlowFromDecimal(d)
}
