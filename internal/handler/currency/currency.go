package currency

import (
	"net/http"

	"github.com/Alitindrawan24/go-currency-conversion/internal/pkg/writer"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) HandleGetCurrency(ctx *gin.Context) {
	currency, err := handler.currencyUseCase.GetCurrency(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
	}

	response := CurrencyResponse{
		Base:     currency.Base,
		Exchange: currency.Exchange,
	}

	ctx.JSON(http.StatusOK, writer.APIResponse("Successfully get currency", "success", response))
}
