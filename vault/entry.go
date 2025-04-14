package vault

import "time"

type Entry struct {
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Vault struct {
	Entries map[string]Entry `json:"entries"`
}
