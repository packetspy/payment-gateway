package service

import (
	"github.com/packetspy/go-payment-gateway/internal/repository"
)

type InvoiceService struct {
	repository repository.InvoiceRepository
}

func NewInvoiceService(repository repository.InvoiceRepository) *InvoiceService {
	return &InvoiceService{repository: repository}
}
