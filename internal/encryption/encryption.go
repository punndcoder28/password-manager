package encryption

import (
	"crypto/rand"
	"crypto/sha256"

	"golang.org/x/crypto/argon2"
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

func KDFGenerator(salt []byte, passKey []byte) []byte {
	keyByte := argon2.IDKey(passKey, salt, 1, 64*1024, 4, 32)
	return keyByte
}
