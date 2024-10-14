package t

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/shopspring/decimal"
)

/*
* 1點為0.0001，即萬分之一
* 在資料庫使用此欄位時，須使用INTEGER
* 在資料庫不使用NUMEIC是因為這樣容量需求比較小(4 bytes)
 */
type Ratio uint32

func (r Ratio) Decimal() decimal.Decimal {
	d := decimal.NewFromUint64(uint64(r)).Shift(-4)
	return d
}

func (r Ratio) MultiplyAmount(amount Amount) Amount {
	result := r.Decimal().Mul(amount.Decimal)
	return Amount{
		Decimal: result,
	}
}
func (r Ratio) Add(y Ratio) Ratio {
	return r + y
}
func (r Ratio) Sub(y Ratio) (Ratio, error) {
	if y > r {
		return 0, fmt.Errorf("Sub Ratio %d - %d result NEGATIVE", r, y)
	}
	return r - y, nil
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

	// user輸入的0.01(%) 代表 1點的Ratio，即萬分之一
	*r = Ratio(d.Shift(2).IntPart())
	return nil
}

// str: 輸入最小單位0.0001，小於0.0001的值將會被忽略
func NewRatioFromString(str string) (Ratio, error) {
	d, err := decimal.NewFromString(str)
	if err != nil {
		return 0, err
	}

	return Ratio(d.Shift(4).IntPart()), nil
}

// 最小單位0.0001，小於0.0001的值將會被忽略
func NewRatioFromDecimal(d decimal.Decimal) (Ratio, error) {
	return Ratio(d.Shift(4).IntPart()), nil
}
