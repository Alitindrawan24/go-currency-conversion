package panichandler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alitindrawan24/go-currency-conversion/internal/pkg/writer"
	"github.com/gin-gonic/gin"
)

// HandlePanic handle panic
func (middleware *Middleware) HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			// Recover the panic if exists
			recover := recover()
			if recover == nil {
				return
			}

			errMsg := fmt.Sprintf("%v", recover)
			ctx.AbortWithStatusJSON(http.StatusBadGateway, writer.APIErrorResponse("Something went wrong", errors.New(errMsg)))
		}()

		ctx.Next()
	}
}
