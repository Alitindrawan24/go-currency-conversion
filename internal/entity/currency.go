package entity

type Currency struct {
	Base     string             `json:"base"`
	Exchange []CurrencyExchange `json:"currency_exchange"`
}

type CurrencyExchange struct {
	Currency string  `json:"currency"`
	Exchange float64 `json:"exchange"`
}
