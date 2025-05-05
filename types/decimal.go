package t

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"

	"github.com/shopspring/decimal"
)

// 差在輸出JSON的時候, 會給Number, 並且可以設定DecimalPoint來指定輸出的數字的小數點, 通常用於API
type DecimalNumber struct {
	decimal.Decimal
	DecimalPoint *int32
}

func (d DecimalNumber) EncodeValues(key string, v *url.Values) error {
	if d.DecimalPoint == nil {
		v.Add(key, d.String())
	} else {
		v.Add(key, d.StringFixed(*d.DecimalPoint))
	}
	return nil
}

func (d *DecimalNumber) MarshalJSON() ([]byte, error) {
	var str string
	if d.DecimalPoint == nil {
		str = d.String()
	} else {
		str = d.StringFixed(*d.DecimalPoint)
	}
	return []byte(str), nil
}

func (dec *DecimalNumber) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return fmt.Errorf("DecimalNumber.UnmarshalXML DecodeElement error: %v", err)
	}
	dc, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	dec.Decimal = dc
	return nil
}

// 差在輸出JSON的時候, 會給string, 通常用於API
type DecimalString struct {
	decimal.Decimal
}

func (d DecimalString) EncodeValues(key string, v *url.Values) error {
	v.Add(key, d.String())
	return nil
}

func (d *DecimalString) MarshalJSON() ([]byte, error) {
	str := d.String()
	return []byte(strconv.Quote(str)), nil
}

func (dec *DecimalString) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return fmt.Errorf("DecimalString.UnmarshalXML DecodeElement error: %v", err)
	}
	dc, err := decimal.NewFromString(s)
	if err != nil {
		return err
	}
	dec.Decimal = dc
	return nil
}
