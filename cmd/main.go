package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"

	"github.com/Alitindrawan24/go-currency-conversion/internal/app"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

func main() {
	// init resty
	client := resty.New()

	// init all layers
	repository := app.NewRepository(client)
	usecase := app.NewUseCase(repository)
	handler := app.NewHandler(usecase)
	middleware := app.NewMiddleware(usecase)

	// init routes
	router := gin.Default()

	// register middleware
	router.Use(middleware.PanicHandler.HandlePanic())

	// GET /
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"app":     os.Getenv("APP_NAME"),
			"version": 0.1,
		})
	})

	// group routes for API
	api := router.Group("/api")

	// GET /currency
	api.GET("/currency", handler.Currency.HandleGetCurrency)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		err := router.Run(fmt.Sprintf(":%s", port))
		if err != nil {
			panic(fmt.Sprintf("Failed to start the web server: %v", err))
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	fmt.Printf("Server stopped\n")
}
