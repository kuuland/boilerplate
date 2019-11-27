package pkg

import (
	"github.com/kuuland/boilerplate/core"
	"github.com/kuuland/boilerplate/pkg/book"
	"github.com/kuuland/boilerplate/pkg/member"
	"github.com/kuuland/kuu"
)

// Mod
func All() *kuu.Mod {
	return &kuu.Mod{
		Code: "bank",
		Models: []interface{}{
			&core.Bank{},
			&core.BankCard{},
			&core.Customer{},
			&core.IDCard{},
			&core.Book{},
			&core.Member{},
		},
		Routes: kuu.RoutesInfo{
			book.Public(),
			book.Private(),
			member.Signup(),
			member.Login(),
		},
	}
}
