package currency

import (
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// Interface of repository contains contract of the repository
type RepositoryInterface interface {
	FetchExchangeByCurrency(*gin.Context, string) (APIResponse, error)
}

// Struct of repository
type Repository struct {
	client *resty.Client
}

func New(client *resty.Client) RepositoryInterface {
	return &Repository{
		client,
	}
}
