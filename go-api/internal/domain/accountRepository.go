package domain

type AccountRepository interface {
	CreateAccount(account *Account) error
	FindByAPIKey(apiKey string) (*Account, error)
	FindByPublicId(publicId string) (*Account, error)
	UpdateBalance(account *Account) error
}
