package c_test

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"math/big"
	"testing"
	"time"

	c "github.com/mychiux413/goutils/common"
)

func TestAES(t *testing.T) {
	h, err := c.RandomHex(32)
	if err != nil {
		t.Error(err)
		return
	}
	a, err := c.NewAESGCM(h)
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 100; i++ {
		dataSize, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			t.Error(err)
			return
		}
		secretData, err := c.RandomBytes(int(dataSize.Int64()))
		if err != nil {
			t.Error(err)
			return
		}

		cipherData, err := a.EncryptBytes(secretData)
		if err != nil {
			t.Error(err)
			return
		}

		decryptedSecretData, err := a.DecryptBytes(cipherData)
		if err != nil {
			t.Error(err)
			return
		}
		if !bytes.Equal(decryptedSecretData, secretData) {
			t.Errorf("expect decryptedSecret to be `%x` but got `%x`", secretData, decryptedSecretData)
		}

		before := time.Now().Add(-time.Second)
		cipherData, err = a.EncryptBytesWithExpired(secretData, before)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = a.DecryptBytesWithExpired(cipherData)
		if !errors.Is(err, c.ErrTokenExpired) {
			t.Errorf("expect expired, but got: %v", err)
			return
		}

		after := time.Now().Add(time.Second)
		cipherData, err = a.EncryptBytesWithExpired(secretData, after)
		if err != nil {
			t.Error(err)
			return
		}
		decryptedSecretData, err = a.DecryptBytesWithExpired(cipherData)
		if err != nil {
			t.Error("expect not expired, but got nil")
			return
		}
		if !bytes.Equal(decryptedSecretData, secretData) {
			t.Errorf("expect decryptedSecretWithExpired to be `%x` but got `%x`", secretData, decryptedSecretData)
		}

		secretHex := hex.EncodeToString(secretData)

		cipherSecret, err := a.EncryptString(secretHex)
		if err != nil {
			t.Error(err)
			return
		}

		decryptedSecretHex, err := a.DecryptString(cipherSecret)
		if err != nil {
			t.Error(err)
			return
		}
		if decryptedSecretHex != secretHex {
			t.Errorf("expect decryptedSecretHex to be `%s` but got `%s`", cipherSecret, secretHex)
			return
		}

		before = time.Now().Add(-time.Second)
		cipherSecret, err = a.EncryptStringWithExpired(secretHex, before)
		if err != nil {
			t.Error(err)
			return
		}
		_, err = a.DecryptStringWithExpired(cipherSecret)
		if !errors.Is(err, c.ErrTokenExpired) {
			t.Errorf("expect expired, but got: %v", err)
			return
		}

		after = time.Now().Add(time.Second)
		cipherSecret, err = a.EncryptStringWithExpired(secretHex, after)
		if err != nil {
			t.Error(err)
			return
		}
		decryptedSecretHex, err = a.DecryptStringWithExpired(cipherSecret)
		if err != nil {
			t.Error("expect not expired, but got nil")
			return
		}
		if decryptedSecretHex != secretHex {
			t.Errorf("expect decryptedSecretHex to be `%s` but got `%s`", cipherSecret, secretHex)
			return
		}

	}
}