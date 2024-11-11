package c

import (
	"errors"
	"regexp"
)

var alphaReg *regexp.Regexp = regexp.MustCompile(`[a-zA-Z]+`)
var digitsReg *regexp.Regexp = regexp.MustCompile(`[0-9]+`)

func ValidatePasswordFormat(password string) error {
	if !alphaReg.MatchString(password) {
		return errors.New("需包含英文")
	}
	if !digitsReg.MatchString(password) {
		return errors.New("需包含數字")
	}
	if len(password) < 6 {
		return errors.New("密碼長度需大於6")
	}
	return nil
}
