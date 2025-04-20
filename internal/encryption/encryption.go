package encryption

import "crypto/sha256"

type Encryptor struct {
	key []byte
}

func NewEncryptor(passkey string) *Encryptor {
	hash := sha256.Sum256([]byte(passkey))
	return &Encryptor{
		key: hash[:],
	}
}
