package app

import "github.com/Alitindrawan24/go-currency-conversion/internal/handler/currency"

type Handlers struct {
	Currency *currency.Handler
}

func NewHandler(usecase *UseCases) *Handlers {
	return &Handlers{
		Currency: currency.New(usecase.Currency),
	}
}
