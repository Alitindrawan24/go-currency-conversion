package app

import (
	"github.com/Alitindrawan24/go-currency-conversion/internal/repository/api/currency"
	"github.com/go-resty/resty/v2"
)

type Repositories struct {
	currency currency.RepositoryInterface
}

func NewRepository(client *resty.Client) *Repositories {
	return &Repositories{
		currency: currency.New(client),
	}
}
