package decrypt

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"crypter/pkg/crypter/encrypt"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Config struct {
	Key   string
	Value string
}

const base64regexp = "(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?"

func Decrypt(ctx context.Context, out io.Writer, cfg Config) error {
	key, err := convertStringToPrivateKey(cfg.Key)
	if err != nil {
		return fmt.Errorf("failed to convert string to private key. %w", err)
	}

	re, err := regexp.Compile(encrypt.EncryptionTemplate + base64regexp + encrypt.EncryptionTemplate)
	if err != nil {
		return fmt.Errorf("compile regexp patter. %w", err)
	}

	result := cfg.Value

	matches := re.FindAllString(result, -1)
	for _, match := range matches {
		value, err := decode(match, key)
		if err != nil {
			return fmt.Errorf("failed decode value with private key. %w", err)
		}
		result = strings.Replace(result, match, value, 1)
	}

	_, err = out.Write([]byte(result))
	return err
}

func decode(s string, key *rsa.PrivateKey) (string, error) {
	s = strings.Replace(s, encrypt.EncryptionTemplate, "", -1)

	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("failed decode string from base64. %w", err)
	}

	decrypted, err := rsa.DecryptOAEP(sha1.New(), rand.Reader, key, b, nil)
	if err != nil {
		return "", fmt.Errorf("failed decrypt value with private key. %w", err)
	}
	return string(decrypted), nil
}

func convertStringToPrivateKey(s string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(s))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the key")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return key, nil
}
