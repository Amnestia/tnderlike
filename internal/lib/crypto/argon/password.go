package argon

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/amnestia/tnderlike/pkg/logger"
	"golang.org/x/crypto/argon2"
)

type config struct {
	memory    uint32
	iteration uint32
	parallel  uint8
	keyLen    uint32

	version int
	format  string
}

func generateRandomSalt() ([]byte, error) {
	b := make([]byte, 256)
	_, err := rand.Read(b)
	if err != nil {
		return b, logger.ErrorWrap(err, "lib", "Error on generating salt")
	}
	return b, nil
}

// GenerateHash generate password hash
func GenerateHash(str, pepper string) (string, error) {
	argonConf := config{
		memory:    64 * 1024,
		iteration: 1,
		parallel:  1,
		keyLen:    256,
		format:    "$argon2id$v=%v$m=%v,t=%v,p=%v$%v$%v",
		version:   argon2.Version,
	}

	salt, err := generateRandomSalt()
	if err != nil {
		return "", err
	}
	pep, err := base64.URLEncoding.DecodeString(pepper)
	if err != nil {
		return "", err
	}
	salt = append(salt, pep...)
	result := argon2.IDKey([]byte(str), salt, argonConf.iteration, argonConf.memory, argonConf.parallel, argonConf.keyLen)

	res := base64.StdEncoding.EncodeToString(result)
	salted := base64.StdEncoding.EncodeToString(salt)
	return fmt.Sprintf(argonConf.format,
		argonConf.version,
		argonConf.memory,
		argonConf.iteration,
		argonConf.parallel,
		salted, res), nil
}

// VerifyHash verify password hash
func VerifyHash(pw, hash string) (bool, error) {
	var (
		conf   = config{}
		err    error
		salt   []byte
		hashed []byte
	)
	hashArr := strings.Split(hash, "$")
	if len(hashArr) < 6 {
		return false, logger.ErrorWrap(fmt.Errorf("invalid hash value"), "util", "Failed on getting hash")
	}

	_, err = fmt.Sscanf(hashArr[2], "v=%d", &conf.version)
	if err != nil {
		return false, logger.ErrorWrap(fmt.Errorf("invalid hash value"), "util", "Failed on getting hash")
	}

	_, err = fmt.Sscanf(hashArr[3], "m=%d,t=%d,p=%d", &conf.memory, &conf.iteration, &conf.parallel)
	if err != nil {
		return false, logger.ErrorWrap(fmt.Errorf("invalid hash value"), "util", "Failed on getting hash")
	}

	salt, err = base64.StdEncoding.DecodeString(hashArr[4])
	if err != nil {
		return false, logger.ErrorWrap(fmt.Errorf("invalid hash value"), "util", "Failed on getting hash")
	}
	hashed, err = base64.StdEncoding.DecodeString(hashArr[5])
	if err != nil {
		return false, logger.ErrorWrap(fmt.Errorf("invalid hash value"), "util", "Failed on getting hash")
	}
	conf.keyLen = uint32(len(hashed))

	result := argon2.IDKey([]byte(pw), salt, conf.iteration, conf.memory, conf.parallel, conf.keyLen)
	return subtle.ConstantTimeCompare(hashed, result) == 1, nil
}
