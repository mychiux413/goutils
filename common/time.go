package c

import (
	"fmt"
	"time"
)

func PointerNow() *time.Time {
	return PointerTime(time.Now())
}

func MustIntervalError(from time.Time, to time.Time) error {
	if from.IsZero() {
		return fmt.Errorf("[%w] from is zero", ErrTimeIntervalError)
	}
	if to.IsZero() {
		return fmt.Errorf("[%w] to is zero", ErrTimeIntervalError)
	}
	if !to.After(from) {
		return fmt.Errorf("[%w] to (%v) must after from (%v)", ErrTimeIntervalError, to, from)
	}
	return nil
}
