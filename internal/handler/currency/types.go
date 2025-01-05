package currency

import "github.com/Alitindrawan24/go-currency-conversion/internal/entity"

type CurrencyResponse struct {
	Base     string                    `json:"base"`
	Exchange []entity.CurrencyExchange `json:"exchange"`
}
