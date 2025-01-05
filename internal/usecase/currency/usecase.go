package currency

import (
	"github.com/Alitindrawan24/go-currency-conversion/internal/entity"
	"github.com/Alitindrawan24/go-currency-conversion/internal/repository/api/currency"
	"github.com/gin-gonic/gin"
)

// Interface of usecase contains contract of the use case
type UseCaseInterface interface {
	GetCurrency(*gin.Context) (entity.Currency, error)
	UpdateCurrency(*gin.Context) error
}

type UseCase struct {
	currencyRepository currency.RepositoryInterface
}

// New initializes usecase.
func New(currencyRepository currency.RepositoryInterface) UseCaseInterface {
	return &UseCase{
		currencyRepository,
	}
}
