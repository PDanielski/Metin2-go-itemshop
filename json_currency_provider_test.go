package main

import (
	"mt2is/pkg/wallet"
	"testing"
)

func TestJsonCurrencyProvider(t *testing.T) {
	conf := []byte(`
	{
		"currency1" : {
			"name": "Currency1 name",
			"iconLink": "currency1://link"
		},
		"currency2" : {
			"name": "Currency2 name",
			"iconLink": "currency2://link"
		}
	}`)

	var expectedCurrencies [2](*wallet.Currency)
	expectedCurrencies[0] = wallet.NewCurrency("currency1", "Currency1 name", "currency1://link")
	expectedCurrencies[1] = wallet.NewCurrency("currency2", "Currency2 name", "currency2://link")
	eq := func(a, b *wallet.Currency) bool {
		return a.Code == b.Code && a.IconLink == b.IconLink && a.Name == b.Name
	}
	in := func(needle *wallet.Currency, heystack []*wallet.Currency) bool {
		for _, currency := range heystack {
			if eq(currency, needle) {
				return true
			}
		}
		return false
	}
	provider, _ := NewJsonCurrencyProvider(conf)
	currencies, _ := provider.Provide()
	for _, c := range currencies {
		if !in(c, expectedCurrencies[:]) {
			t.Errorf("%+v is not expected\n", c)
		}
	}

}
