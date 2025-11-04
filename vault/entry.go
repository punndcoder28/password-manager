package vault

import "time"

type Entry struct {
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	IsActive      bool      `json:"is_active"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeactivatedAt time.Time `json:"deactivated_at"`
}

type Vault struct {
	Entries map[string][]Entry `json:"entries"`
}

type MaskedEntry struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
