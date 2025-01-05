package currency

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func (repository *Repository) FetchExchangeByCurrency(ctx *gin.Context, currency string) (APIResponse, error) {
	result := APIResponse{}
	resp, err := repository.client.R().
		SetResult(&result).
		SetQueryParams(map[string]string{
			"apikey":        os.Getenv("FREE_CURRENCY_API_KEY"),
			"base_currency": currency,
		}).
		Get("https://api.freecurrencyapi.com/v1/latest")

	fmt.Println(resp.StatusCode())

	return result, err
}
