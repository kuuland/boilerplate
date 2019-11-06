package bank

import (
	"github.com/kuuland/kuu"
)

// Mod
func Mod() *kuu.Mod {
	return &kuu.Mod{
		Code: "bank",
		Models: []interface{}{
			&Bank{},
			&BankCard{},
			&Customer{},
			&IDCard{},
		},
	}
}
