package domain

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	PublicId  string
	Name      string
	Email     string
	APIKey    string
	Balance   float64
	Mu        sync.RWMutex
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GenerateAPIKey() string {
	key := make([]byte, 16)
	rand.Read(key)
	return hex.EncodeToString(key)
}

func NewAccount(name, email string) *Account {
	account := &Account{
		PublicId:  uuid.New().String(),
		Name:      name,
		Email:     email,
		APIKey:    GenerateAPIKey(),
		Balance:   0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return account
}

func (a *Account) AddBalance(amount float64) {
	a.Mu.Lock()
	defer a.Mu.Unlock()
	a.Balance += amount
	a.UpdatedAt = time.Now()
}
