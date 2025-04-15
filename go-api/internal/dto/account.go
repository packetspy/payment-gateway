package dto

import (
	"time"

	"github.com/packetspy/go-payment-gateway/internal/domain"
)

type CreateAccountRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AccountResponse struct {
	PublicId  string    `json:"publicId"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   float64   `json:"balance"`
	ApiKey    string    `json:"apiKey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ToAccount(input CreateAccountRequest) *domain.Account {
	return domain.NewAccount(input.Name, input.Email)
}

func FromAccount(account *domain.Account) AccountResponse {
	return AccountResponse{
		PublicId:  account.PublicId,
		Name:      account.Name,
		Email:     account.Email,
		Balance:   account.Balance,
		ApiKey:    account.APIKey,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}
}
