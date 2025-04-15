package service

import (
	"github.com/packetspy/go-payment-gateway/internal/domain"
	"github.com/packetspy/go-payment-gateway/internal/dto"
	"github.com/packetspy/go-payment-gateway/internal/repository"
)

type AccountService struct {
	repository repository.AccountRepository
}

func NewAccountService(repository *repository.AccountRepository) *AccountService {
	return &AccountService{repository: *repository}
}

func (s *AccountService) CreateAccount(input dto.CreateAccountRequest) (*dto.AccountResponse, error) {
	account := dto.ToAccount(input)
	existingAccount, err := s.repository.FindByAPIKey(account.APIKey)
	if err != nil && err != domain.ErrorAccountNotFound {
		return nil, err
	}

	if existingAccount != nil {
		return nil, domain.ErrorAccountAlreadyExists
	}

	err = s.repository.CreateAccount(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByAPIKey(apiKey string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) FindByPublicId(publicId string) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByPublicId(publicId)
	if err != nil {
		return nil, err
	}
	output := dto.FromAccount(account)
	return &output, nil
}

func (s *AccountService) UpdateBalance(apiKey string, amount float64) (*dto.AccountResponse, error) {
	account, err := s.repository.FindByAPIKey(apiKey)
	if err != nil {
		return nil, err
	}

	account.AddBalance(amount)
	err = s.repository.UpdateBalance(account)
	if err != nil {
		return nil, err
	}

	output := dto.FromAccount(account)
	return &output, nil
}
