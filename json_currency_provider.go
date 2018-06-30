package main

import (
	"encoding/json"
	"fmt"
	"mt2is/pkg/wallet"
)

func NewJsonCurrencyProvider(j []byte) (wallet.CurrencyProvider, error) {
	c := make(map[string](map[string]string))
	err := json.Unmarshal(j, &c)
	if err != nil {
		return nil, fmt.Errorf("Failed to parsing json currency config: %v", err)
	}
	currencies := make([]*wallet.Currency, 0)
	for currencyCode, data := range c {
		currencies = append(currencies, wallet.NewCurrency(currencyCode, data["name"], data["iconLink"]))
	}
	return wallet.NewSimpleCurrencyProvider(currencies), nil
}
