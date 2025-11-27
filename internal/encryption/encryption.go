package encryption

import (
	"crypto/rand"
	"crypto/sha256"
)

type Encryptor struct {
	key []byte
}

func NewEncryptor(passkey string) *Encryptor {
	hash := sha256.Sum256([]byte(passkey))
	return &Encryptor{
		key: hash[:],
	}
}

func GenerateSalt() []byte {
	salt := make([]byte, 32)
	if _, err := rand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}
