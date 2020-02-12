package encrypt

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

const publicKey = `
-----BEGIN RSA PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAMWHvBCOUAJsj11lMYBbc4es8iqSkNAI
qSvPNqBJop5ukSpi3AckngvIu9XLm73TF6XgqgXEv5ye44gyHwGdBs0CAwEAAQ==
-----END RSA PUBLIC KEY-----
`

func Test_Sync(t *testing.T) {
	ctx := context.Background()
	out := logrus.StandardLogger().Out

	cfg := Config{
		Key:   publicKey,
		Value: "long-long-long-word",
	}

	if err := Encrypt(ctx, out, cfg); err != nil {
		t.Fatalf("%+v", err)
	}
}
