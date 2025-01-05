package app

import "github.com/Alitindrawan24/go-currency-conversion/internal/middleware/panichandler"

type Middleware struct {
	PanicHandler *panichandler.Middleware
}

func NewMiddleware(useCase *UseCases) *Middleware {
	return &Middleware{
		PanicHandler: panichandler.New(),
	}
}
