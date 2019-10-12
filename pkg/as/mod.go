package as

import (
	"github.com/kuuland/kuu"
)

func Mod() *kuu.Mod {
	return &kuu.Mod{
		Code: "as",
		Models: []interface{}{
			&Bank{},
			&BankCard{},
			&Customer{},
			&IDCard{},
		},
	}
}
