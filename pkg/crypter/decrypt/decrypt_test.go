package decrypt

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

const privateKey = `
-----BEGIN RSA PRIVATE KEY-----
MIIBPAIBAAJBAMWHvBCOUAJsj11lMYBbc4es8iqSkNAIqSvPNqBJop5ukSpi3Ack
ngvIu9XLm73TF6XgqgXEv5ye44gyHwGdBs0CAwEAAQJAZOmgKXEa7PIbF+KftGyE
DBdNrHQuKSmTi38T8DVOL9N02zns7f8GY9breB3ELq8wPkcTD2+/aISeEbm+5Nva
QQIhANr438xCLMFCkm0RWv51bEgRlvHL2pQ3xHzWBJ/L9VEdAiEA5u6n6DdyRiSc
kIv0xoHd9/pSt29NLDmwx177S5vgzXECIQCLnz00RM28vPIY0YQv1DejDHQu4UkS
USzcXKq+KZLWkQIhAN4jGc6tbzX7x8LfbeB5UcxUtbaP0NtGzz6ope/QDMlxAiEA
svX0pvP04QpNAv1pm6o3R7ry8NTY235H6HdYP5pgJJw=
-----END RSA PRIVATE KEY-----
`

const encodedString = `
email:
  password: '{cipher}eLE7raf9NSO2N/oKYRmJkr8hJfOM+PZaD5yRlXgw7SF4EaSALLSU0FXvrz1y0zyh1lPl6Ko3iLOzmSEjZBC8aQ=={cipher}'
  same-as-above: '{cipher}eLE7raf9NSO2N/oKYRmJkr8hJfOM+PZaD5yRlXgw7SF4EaSALLSU0FXvrz1y0zyh1lPl6Ko3iLOzmSEjZBC8aQ=={cipher}'
mongo:
  user: user
  password: '{cipher}mMsq7F8Ek5Eg/hbyY3P4wXMAaCnToYqbUdbUvKQQiHxblqI2yq9+UYvxGBhCihrwLTTpqep6ImfALT6aRHJiqw=={cipher}'
and-last-one: '{cipher}UMh6zzAGN/NXyXjFlkZv6uM+cllRPnWQ1U7gRmPkmajiGxSmDQW3cN+ISAVFiUDkfigXSQMW3Ot64Tr7Nbdg5g=={cipher}'
`

func Test_Sync(t *testing.T) {
	ctx := context.Background()
	out := logrus.StandardLogger().Out

	cfg := Config{
		Key:   privateKey,
		Value: encodedString,
	}

	if err := Decrypt(ctx, out, cfg); err != nil {
		t.Fatalf("%+v", err)
	}
}
