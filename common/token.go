package c

import (
	"sync"
	"time"
)

type Token struct {
	sync.Mutex
	token     string
	expiredAt time.Time
}

func (t *Token) Get() string {
	t.Lock()
	defer t.Unlock()
	return t.token
}

func (t *Token) Set(token string, expiredAt time.Time) {
	t.Lock()
	defer t.Unlock()
	t.token = token
	t.expiredAt = expiredAt
}

func (t *Token) IsExpired() bool {
	return t.expiredAt.Before(time.Now())
}
