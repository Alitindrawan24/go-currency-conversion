package currency

import (
	"github.com/Alitindrawan24/go-currency-conversion/internal/usecase/currency"
)

type Handler struct {
	currencyUseCase currency.UseCaseInterface
}

func New(currencyUseCase currency.UseCaseInterface) *Handler {
	return &Handler{
		currencyUseCase,
	}
}
