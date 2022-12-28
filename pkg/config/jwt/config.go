package jwtConfigTools

import (
	errorTools "golang-ecommerce-example/pkg/error"
	"os"
)

func LoadSigningKey() ([]byte, error) {
	k := []byte(os.Getenv("SIGNING_KEY"))
	if k == nil {
		return k, errorTools.ErrLoadSigningKey.Error
	}
	return k, nil
}
