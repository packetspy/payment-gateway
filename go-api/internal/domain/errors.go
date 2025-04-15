package domain

import "errors"

var (
	ErrorAccountNotFound      = errors.New("account not found")
	ErrorDuplicatedKey        = errors.New("duplicated key")
	ErrorInvoiceNotFound      = errors.New("invoice not found")
	ErrorUnauthorizedAccess   = errors.New("unauthorized access")
	ErrorInvalidRequest       = errors.New("invalid request")
	ErrorAccountAlreadyExists = errors.New("account already exists")
)
