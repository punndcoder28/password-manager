package passkey

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/cryptobyte"
)

const currentVersion = 1
const memory = 64 * 1024
const iterations = 3
const parallelism = 2
const saltLength = 32
const keyLength = 64

type PasskeyData struct {
	Version   uint32
	Salt      []byte
	HashedKey []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PasskeyManager struct {
	filePath string
	data     *PasskeyData
}

func NewPasskeyManager(configDir string) (*PasskeyManager, error) {
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create config directory: %w", err)
	}

	filePath := filepath.Join(configDir, "passkey.dat")
	return &PasskeyManager{
		filePath: filePath,
	}, nil
}

func (pm *PasskeyManager) InitializePasskey(passkey string) error {
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return fmt.Errorf("failed to generate salt: %w", err)
	}

	hashedKey := hashPasskey(passkey, salt)

	now := time.Now()
	pm.data = &PasskeyData{
		Version:   currentVersion,
		Salt:      salt,
		HashedKey: hashedKey,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := pm.save(); err != nil {
		return fmt.Errorf("failed to save passkey data: %w", err)
	}

	return nil
}

func hashPasskey(passkey string, salt []byte) []byte {
	return argon2.Key([]byte(passkey), salt, iterations, memory, parallelism, keyLength)
}

func (pm *PasskeyManager) save() error {
	if pm.data == nil {
		return fmt.Errorf("no passkey data to save")
	}

	b := cryptobyte.NewBuilder(nil)
	b.AddUint32(pm.data.Version)
	b.AddBytes(pm.data.Salt)
	b.AddBytes(pm.data.HashedKey)

	data, err := b.Bytes()
	if err != nil {
		return fmt.Errorf("failed to build passkey data: %w", err)
	}

	tempFile := pm.filePath + ".tmp"
	if err := os.WriteFile(tempFile, data, 0600); err != nil {
		return fmt.Errorf("failed to write passkey data: %w", err)
	}

	if err := os.Rename(tempFile, pm.filePath); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

func (pm *PasskeyManager) VerifyPasskey(passkey string) (bool, error) {
	if err := pm.load(); err != nil {
		return false, err
	}

	hashedKey := hashPasskey(passkey, pm.data.Salt)
	return secureCompare(hashedKey, pm.data.HashedKey), nil
}

// TODO: understand if we need update passkey function because
// right now it makes no sense because if user has lost the passkey
// then they have lost their passwords as well

func (pm *PasskeyManager) DeriveKey(passkey string) ([]byte, error) {
	if err := pm.load(); err != nil {
		return nil, err
	}

	key := argon2.Key([]byte(passkey), pm.data.Salt, iterations, memory, parallelism, keyLength)
	return key, nil
}

func (pm *PasskeyManager) load() error {
	if pm.data != nil {
		return nil
	}

	data, err := os.ReadFile(pm.filePath)
	if err != nil {
		return fmt.Errorf("failed to read passkey data: %w", err)
	}

	// understand what this is doing and why we need it
	if len(data) < 4 {
		return fmt.Errorf("invalid passkey file format")
	}
	version := binary.BigEndian.Uint32(data[:4])

	if len(data) < 4+saltLength+keyLength {
		return fmt.Errorf("invalid passkey file format")
	}

	pm.data = &PasskeyData{
		Version:   version,
		Salt:      data[4 : 4+saltLength],
		HashedKey: data[4+saltLength : 4+saltLength+keyLength],
	}

	return nil
}

func secureCompare(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	var result byte
	for i := range a {
		result |= a[i] ^ b[i]
	}

	return result == 0
}
