package card

import "github.com/phonkluver/bank/pkg/bank/types"

func Issue(currency types.Currency, color string, name string) types.Card {
	return types.Card{
		ID:       1000,
		PAN:      "5058 XXXX XXXX 0001",
		Balance:  0,
		Currenct: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
}
