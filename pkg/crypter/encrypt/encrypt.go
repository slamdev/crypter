package encrypt

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
)

type Config struct {
	Key   string
	Value string
}

const EncryptionTemplate = "{cipher}"

func Encrypt(ctx context.Context, out io.Writer, cfg Config) error {
	key, err := convertStringToPublicKey(cfg.Key)
	if err != nil {
		return fmt.Errorf("failed to convert string to public key. %w", err)
	}
	encrypted, err := rsa.EncryptOAEP(sha1.New(), rand.Reader, key, bytes.NewBufferString(cfg.Value).Bytes(), nil)
	if err != nil {
		return fmt.Errorf("failed to encrypt value. %w", err)
	}
	result := base64.StdEncoding.EncodeToString(encrypted)
	result = EncryptionTemplate + result + EncryptionTemplate
	_, err = out.Write([]byte(result))
	return err
}

func convertStringToPublicKey(s string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		return nil, errors.New("key type is not RSA")
	}
}
