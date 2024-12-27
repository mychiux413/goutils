package c

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

type AESConfig struct {
	AESBase64Key string `json:"aes_base64_key" yaml:"aes_base64_key"`
	AESHexKey    string `json:"aes_hex_key" yaml:"aes_hex_key"`
}

type AESGCM struct {
	aead cipher.AEAD
}

func NewAESGCMWithBase64(keyInBase64 string) (*AESGCM, error) {
	key, err := base64.RawStdEncoding.DecodeString(keyInBase64)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AESGCM{
		aead: aesGCM,
	}, nil
}
func NewAESGCMWithHex(keyInHex string) (*AESGCM, error) {
	key, err := hex.DecodeString(keyInHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return &AESGCM{
		aead: aesGCM,
	}, nil
}

func (conf *AESConfig) NewAES() (*AESGCM, error) {
	if conf.AESBase64Key != "" {
		return NewAESGCMWithBase64(conf.AESBase64Key)
	}
	return NewAESGCMWithHex(conf.AESHexKey)
}

func (a *AESGCM) getNonce() ([]byte, error) {
	nonce := make([]byte, a.aead.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func (a *AESGCM) EncryptBytes(data []byte) ([]byte, error) {
	nonce, err := a.getNonce()
	if err != nil {
		return nil, err
	}
	return a.aead.Seal(nonce, nonce, data, nil), nil
}

func (a *AESGCM) EncryptString(text string) (string, error) {
	output, err := a.EncryptBytes([]byte(text))
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(output), nil
}

func (a *AESGCM) DecryptBytes(ciphertext []byte) ([]byte, error) {
	nonceSize := a.aead.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, fmt.Errorf("invalid cipher data size: %d", len(ciphertext))
	}
	output, err := a.aead.Open(nil, ciphertext[:nonceSize], ciphertext[nonceSize:], nil)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (a *AESGCM) DecryptString(cipherText string) (string, error) {
	cipherData, err := base64.RawStdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", err
	}
	data, err := a.DecryptBytes(cipherData)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// 將plaintext加密，但是給上過期時間，超過時間了解開就報錯
func (a *AESGCM) EncryptBytesWithExpired(plaintext []byte, expiredAt time.Time) ([]byte, error) {
	timestamp := expiredAt.UnixMicro()
	if timestamp < 0 {
		return nil, fmt.Errorf("expiredAt %v's timestamp: %d < 0", expiredAt, timestamp)
	}
	timestampInBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestampInBytes, uint64(timestamp))

	plaintextWithExpired := append(timestampInBytes, plaintext...)

	return a.EncryptBytes(plaintextWithExpired)
}

func (a *AESGCM) DecryptBytesWithExpired(ciphertext []byte) ([]byte, error) {
	plaintextWithExpired, err := a.DecryptBytes(ciphertext)
	if err != nil {
		return nil, err
	}
	if len(plaintextWithExpired) < 8 {
		return nil, fmt.Errorf("invalid plaintextWithExpired size: %d", len(plaintextWithExpired))
	}

	timestampInBytes := plaintextWithExpired[:8]
	timestamp := int64(binary.LittleEndian.Uint64(timestampInBytes))
	if time.Now().UnixMicro()-timestamp > 0 {
		return nil, ErrTokenExpired
	}
	plaintext := plaintextWithExpired[8:]

	return plaintext, nil
}

func (a *AESGCM) EncryptStringWithExpired(plainstring string, expiredAt time.Time) (string, error) {
	ciphertext, err := a.EncryptBytesWithExpired([]byte(plainstring), expiredAt)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(ciphertext), nil

}

func (a *AESGCM) DecryptStringWithExpired(cipherString string) (string, error) {
	ciphertext, err := base64.RawStdEncoding.DecodeString(cipherString)
	if err != nil {
		return "", err
	}

	plaintext, err := a.DecryptBytesWithExpired(ciphertext)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func RandomBytes(byteSize int) ([]byte, error) {
	bytes := make([]byte, byteSize) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

func RandomHex(byteSize int) (string, error) {
	bytes, err := RandomBytes(byteSize)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
func RandomBase64(byteSize int) (string, error) {
	bytes, err := RandomBytes(byteSize)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(bytes), nil
}
