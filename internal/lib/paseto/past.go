package paseto

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"
	"time"

	"github.com/amnestia/tnderlike/internal/config"
	"github.com/o1egl/paseto/v2"
)

// Handler interface
type Handler interface {
	Generate(Payload) (string, error)
	Extract(string) (Payload, error)
}

// PASTHandle struct
type PASTHandle struct {
	Paseto *paseto.V2
	Config config.Config
}

// New initialize paseto generator
func New(cfg config.Config) *PASTHandle {
	p := paseto.NewV2()
	return &PASTHandle{
		Paseto: p,
		Config: cfg,
	}
}

// Generate generate paseto token
func (h *PASTHandle) Generate(payload Payload) (string, error) {
	if h.Config.Environment == "production" || h.Config.Environment == "staging" {
		return h.generatePublic(payload)
	}
	return h.generateLocal(payload)
}

// Extract extract paseto token
func (h *PASTHandle) Extract(token string) (Payload, error) {
	if h.Config.Environment == "production" || h.Config.Environment == "staging" {
		return h.extractPublic(token)
	}
	return h.extractLocal(token)
}

func (h *PASTHandle) generateClaims(payload Payload) paseto.JSONToken {
	now := time.Now()
	exp := now.Add(time.Minute * 5)
	if payload.TokenType == RefreshToken {
		exp = now.Add(time.Hour * 24 * 3)
	}
	c := paseto.JSONToken{
		Audience:   audience,
		Issuer:     issuer,
		Jti:        jti,
		IssuedAt:   now,
		Expiration: exp,
		NotBefore:  now,
	}
	c.Set(payloadKey, payload)
	c.Set(tokenTypeKey, payload.TokenType)
	return c
}

func (h *PASTHandle) checkClaims(claims paseto.JSONToken) error {
	tokenErr := errors.New("invalid token")
	if claims.Audience != audience {
		return tokenErr
	}
	if claims.Issuer != issuer {
		return tokenErr
	}
	if claims.Jti != jti {
		return tokenErr
	}
	return nil
}

// generatePublic generate paseto token for public mode
func (h *PASTHandle) generatePublic(payload Payload) (string, error) {
	c := h.generateClaims(payload)
	key, err := hex.DecodeString(h.Config.Auth.PrivKey)
	if err != nil {
		return "", err
	}
	privKey := ed25519.PrivateKey(key)
	token, err := h.Paseto.Sign(privKey, c, footer)
	if err != nil {
		return "", err
	}
	return token, nil
}

// extractPublic decrypt paseto token for local mode
func (h *PASTHandle) extractPublic(token string) (ret Payload, err error) {
	key, err := hex.DecodeString(h.Config.Auth.PubKey)
	if err != nil {
		return
	}
	pubKey := ed25519.PublicKey(key)
	claims := paseto.JSONToken{}
	footer := ""
	err = h.Paseto.Verify(token, pubKey, &claims, &footer)
	if err != nil {
		return
	}
	err = claims.Validate()
	if err != nil {
		return
	}
	err = h.checkClaims(claims)
	if err != nil {
		return
	}
	err = claims.Get(payloadKey, &ret)
	if err != nil {
		return
	}
	return
}

// generateLocal generate paseto token for local mode
func (h *PASTHandle) generateLocal(payload Payload) (string, error) {
	c := h.generateClaims(payload)
	key, err := hex.DecodeString(h.Config.Auth.LocalKey)
	if err != nil {
		return "", err
	}
	token, err := h.Paseto.Encrypt(key, c, footer)
	if err != nil {
		return "", err
	}
	return token, nil
}

// extractLocal decrypt paseto token for local mode
func (h *PASTHandle) extractLocal(token string) (ret Payload, err error) {
	key, err := hex.DecodeString(h.Config.Auth.LocalKey)
	if err != nil {
		return
	}
	claims := paseto.JSONToken{}
	footer := ""
	err = h.Paseto.Decrypt(token, key, &claims, &footer)
	if err != nil {
		return
	}
	err = claims.Validate()
	if err != nil {
		return
	}
	err = h.checkClaims(claims)
	if err != nil {
		return
	}
	err = claims.Get(payloadKey, &ret)
	if err != nil {
		return
	}
	return
}
