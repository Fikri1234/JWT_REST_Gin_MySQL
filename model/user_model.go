package model

// MUser struct
type MUser struct {
	ID                 int64  `json:"id"`
	UserName           string `json:"userName"`
	Password           string `json:"password"`
	AccountExpired     bool   `json:"accountExpired"`
	AccountLocked      bool   `json:"accountLocked"`
	CredentialsExpired bool   `json:"credentialsExpired"`
	Enabled            bool   `json:"enabled"`
}

// MUsers array of MUser type
type MUsers []MUser
