package app

import "github.com/Alitindrawan24/go-currency-conversion/internal/usecase/currency"

type UseCases struct {
	Currency currency.UseCaseInterface
}

func NewUseCase(repository *Repositories) *UseCases {
	return &UseCases{
		Currency: currency.New(repository.currency),
	}
}
