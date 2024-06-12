package authsvc

import "github.com/amnestia/tnderlike/internal/lib/crypto/argon"

var (
	handleGenerateHash = func(s string, pepper string) (string, error) {
		return argon.GenerateHash(s, pepper)
	}

	handleVerifyHash = func(s string, p string) (bool, error) {
		return argon.VerifyHash(s, p)
	}
)
