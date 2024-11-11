package c

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type Secreter interface {
	GetMerchantCode() (string, error)
	GetAccount() (string, error)
	GetPassword() (string, error)
	GetKey1() (string, error)
	GetKey2() (string, error)
	GetKey3() (string, error)
	GetKey4() (string, error)

	// Must Prefix with `https://`
	GetEndpointURL1() (string, error)

	// Must Prefix with `https://`
	GetEndpointURL2() (string, error)
	// Must Prefix with `https://`
	GetEndpointURL3() (string, error)
	// Must Prefix with `https://`
	GetEndpointURL4() (string, error)
}

type RequiredSetting struct {
	MerchantCode bool
	Account      bool
	Password     bool
	Key1         bool
	Key2         bool
	Key3         bool
	Key4         bool
	Endpoint1    bool
	Endpoint2    bool
	Endpoint3    bool
	Endpoint4    bool
}

type SecretRequiredSetting struct {
	Name            string
	RequiredSetting *RequiredSetting
}

type Secret struct {
	MerchantCode string `yaml:"MerchantCode"`
	Account      string `yaml:"Account"`
	Password     string `yaml:"Password"`
	Key1         string `yaml:"Key1"`
	Key2         string `yaml:"Key2"`
	Key3         string `yaml:"Key3"`
	Key4         string `yaml:"Key4"`
	Endpoint1    string `yaml:"Endpoint1"`
	Endpoint2    string `yaml:"Endpoint2"`
	Endpoint3    string `yaml:"Endpoint3"`
	Endpoint4    string `yaml:"Endpoint4"`
}

func (s *Secret) GetMerchantCode() (string, error) {
	if s.MerchantCode == "" {
		return "", errors.New("MerchantCode is empty")
	}
	return s.MerchantCode, nil
}
func (s *Secret) GetAccount() (string, error) {
	if s.Account == "" {
		return "", errors.New("Account is empty")
	}
	return s.Account, nil
}
func (s *Secret) GetPassword() (string, error) {
	if s.Password == "" {
		return "", errors.New("Password is empty")
	}
	return s.Password, nil
}
func (s *Secret) GetKey1() (string, error) {
	if s.Key1 == "" {
		return "", errors.New("Key1 is empty")
	}
	return s.Key1, nil
}
func (s *Secret) GetKey2() (string, error) {
	if s.Key2 == "" {
		return "", errors.New("Key2 is empty")
	}
	return s.Key2, nil
}
func (s *Secret) GetKey3() (string, error) {
	if s.Key3 == "" {
		return "", errors.New("Key3 is empty")
	}
	return s.Key3, nil
}
func (s *Secret) GetKey4() (string, error) {
	if s.Key4 == "" {
		return "", errors.New("Key4 is empty")
	}
	return s.Key4, nil
}
func (s *Secret) GetEndpointURL1() (string, error) {
	return s.Endpoint1, nil
}
func (s *Secret) GetEndpointURL2() (string, error) {
	return s.Endpoint2, nil
}
func (s *Secret) GetEndpointURL3() (string, error) {
	return s.Endpoint3, nil
}
func (s *Secret) GetEndpointURL4() (string, error) {
	return s.Endpoint4, nil
}

type DevSecretMap map[string]Secret

// 儲存Secret的Key
type StoreKey string

func ToStoreKey(prefix, apiModule string) StoreKey {
	return StoreKey(url.QueryEscape(prefix + "-" + apiModule))
}
func (key StoreKey) Split() (prefix string, apiModule string) {
	s := strings.SplitN(key.String(), "-", 2)
	if len(s) != 2 {
		err := fmt.Errorf("must split into [2]string, but got length: %d", len(s))
		panic(err)
	}

	prefix = s[0]
	apiModule = s[1]
	return
}
func (key StoreKey) String() string {
	return string(key)
}
