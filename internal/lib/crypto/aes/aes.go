package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
)

// Encrypt encrypt string and return the encrypted version in base64
func Encrypt(key, target string) (ret string, err error) {
	target = base64.StdEncoding.EncodeToString([]byte(target))
	acKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	aesBlock, err := aes.NewCipher(acKey)
	if err != nil {
		return ret, err
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return ret, err
	}

	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return ret, err
	}

	crypted := gcm.Seal(nonce, nonce, []byte(target), nil)
	ret = base64.StdEncoding.EncodeToString(crypted)

	return
}

// Decrypt decrypt encrypted string and return the plaintext
func Decrypt(key, target string) (ret string, err error) {
	acKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return "", err
	}

	aesBlock, err := aes.NewCipher([]byte(acKey))
	if err != nil {
		return ret, err
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return ret, err
	}

	acTarget, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	nonce, crypted := acTarget[:nonceSize], acTarget[nonceSize:]
	decoded, err := gcm.Open(nil, []byte(nonce), []byte(crypted), nil)
	if err != nil {
		return "", err
	}

	decoded, err = base64.StdEncoding.DecodeString(string(decoded))
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
