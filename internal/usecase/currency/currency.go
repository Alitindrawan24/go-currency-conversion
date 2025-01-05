package currency

import (
	"github.com/Alitindrawan24/go-currency-conversion/internal/entity"
	"github.com/gin-gonic/gin"
)

func (usecase *UseCase) GetCurrency(ctx *gin.Context) (entity.Currency, error) {
	result, err := usecase.currencyRepository.FetchExchangeByCurrency(ctx, "IDR")
	if err != nil {
		return entity.Currency{}, err
	}

	currencyExchange := entity.Currency{
		Base: "IDR",
	}
	for currency, value := range result.Data {
		currencyExchange.Exchange = append(currencyExchange.Exchange, entity.CurrencyExchange{
			Currency: currency,
			Exchange: 1 / value,
		})
	}

	return currencyExchange, nil
}

func (usecase *UseCase) UpdateCurrency(ctx *gin.Context) error {
	return nil
}
