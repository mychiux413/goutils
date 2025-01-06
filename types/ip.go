package t

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"

	"github.com/asaskevich/govalidator"
)

type IP string

func ParseIP(ipString string) IP {

	if !govalidator.IsIP(ipString) {
		ipString = "127.0.0.1"
	}

	return IP(ipString)

}

func (ip IP) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(string(ip))))
}

func (ip *IP) UnmarshalGQL(v interface{}) error {
	value, ok := v.(string)
	if !ok {
		return fmt.Errorf("IP must be a string")
	}
	*ip = IP(value)

	return nil
}

func (i *IP) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("IP should be string, but got %s", reflect.TypeOf(value))
	}
	*i = IP(str)
	return nil
}

func (i IP) Value() (driver.Value, error) {
	if len(i) == 0 {
		return nil, nil
	}
	return string(i), nil
}

type IPArray []IP

func (ips *IPArray) Scan(value interface{}) error {
	strs, err := sqlStrToStrings(value)
	if err != nil {
		return err
	}
	for _, str := range strs {
		var ip IP
		err := ip.Scan(str)
		if err != nil {
			return err
		}
		*ips = append(*ips, ip)
	}
	return nil
}

func (ips IPArray) Value() (driver.Value, error) {
	if len(ips) == 0 {
		return "{}", nil
	}
	var strs []string
	for _, ip := range ips {
		strs = append(strs, string(ip))
	}
	arr := "{" + strings.Join(strs, ",") + "}"
	return arr, nil
}

func (ips *IPArray) UnmarshalGQL(v interface{}) error {
	return errors.New("IPArray.UnmarshalGQL not implemented")
}

func (ips *IPArray) Includes(ip IP) bool {
	for _, i := range *ips {
		if i == ip {
			return true
		}
	}
	return false
}

// 回傳不重複的IP
func (ips *IPArray) Unique() IPArray {
	var output IPArray
	mp := map[IP]bool{}
	for _, ip := range *ips {
		mp[ip] = true
	}
	for id := range mp {
		output = append(output, id)
	}
	return output
}

func UniqueIPs(ips []IP) []IP {
	ipMap := map[IP]bool{}
	for _, ip := range ips {
		if ip == "" {
			continue
		}
		ipMap[ip] = true
	}

	var uniques []IP
	for ip := range ipMap {
		uniques = append(uniques, ip)
	}
	return uniques
}
