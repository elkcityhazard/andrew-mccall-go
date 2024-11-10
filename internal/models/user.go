package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"time"
)

type Password struct {
	ID        int64
	PlainText string
	Hash      string
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
	Version   int64
}

type User struct {
	ID        int64
	Email     string
	Username  string
	Password  *Password
	CreatedAt time.Time
	UpdatedAt time.Time
	IsActive  bool
	Role      string
	Version   int64
}

type ActivationToken struct {
	ID        int64
	Plaintext string
	Hash      []byte
	UserID    int64
	Expiry    time.Time
	Scope     string
}

func NewActivationToken() *ActivationToken {
	return &ActivationToken{}
}

func (at *ActivationToken) GenerateToken(userID int64, ttl time.Duration, scope string) error {
	at.UserID = userID
	at.Expiry = time.Now().Add(ttl)
	at.Scope = scope

	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)

	if err != nil {
		return err
	}

	at.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	hash := sha256.Sum256([]byte(at.Plaintext))
	at.Hash = hash[:]

	return nil

}
