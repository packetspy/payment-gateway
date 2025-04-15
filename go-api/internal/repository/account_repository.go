package repository

import (
	"database/sql"
	"time"

	"github.com/packetspy/go-payment-gateway/internal/domain"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(account *domain.Account) error {
	stmt, err := r.db.Prepare(`
	INSERT INTO accounts (public_id, name, email, api_key, balance, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`)

	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		account.PublicId,
		account.Name,
		account.Email,
		account.APIKey,
		account.Balance,
		account.CreatedAt,
		account.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *AccountRepository) FindByAPIKey(apiKey string) (*domain.Account, error) {
	var account domain.Account
	err := r.db.QueryRow(`
	SELECT public_id, name, email, api_key, balance, created_at, updated_at
	FROM accounts
	WHERE api_key = $1`, apiKey).Scan(
		&account.PublicId,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrorAccountNotFound
	}

	return &account, nil
}

func (r *AccountRepository) FindByPublicId(publicId string) (*domain.Account, error) {
	var account domain.Account
	err := r.db.QueryRow(`
	SELECT public_id, name, email, api_key, balance, created_at, updated_at
	FROM accounts
	WHERE public_id = $1`, publicId).Scan(
		&account.PublicId,
		&account.Name,
		&account.Email,
		&account.APIKey,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, domain.ErrorAccountNotFound
	}

	return &account, nil
}

func (r *AccountRepository) UpdateBalance(account *domain.Account) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`
	SELECT balance FROM accounts WHERE public_id = $1 FOR UPDATE
	`, account.PublicId).Scan(&currentBalance)

	if err == sql.ErrNoRows {
		return domain.ErrorAccountNotFound
	}

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
	UPDATE accounts
	SET balance = $1, updated_at = $2
	WHERE public_id = $3
	`, account.Balance, time.Now(), account.PublicId)

	if err != nil {
		return err
	}

	return tx.Commit()
}
