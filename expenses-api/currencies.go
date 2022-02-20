package main

import "fmt"

var currencies map[string]Currency
var loaded bool

type Currency struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func getCurrencies() map[string]Currency {
	if !loaded {
		loadCurrencies()
	}
	return currencies
}

func loadCurrencies() {
	currencies = make(map[string]Currency)
	currencies["BTC"] = Currency{"Bitcoin", "BTC"}
	currencies["USD"] = Currency{"United States dollar", "USD"}
	currencies["EUR"] = Currency{"Euro", "EUR"}
	loaded = true
	fmt.Println("Currencies loaded")
}
