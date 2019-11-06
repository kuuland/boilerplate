package bank

import (
	"github.com/kuuland/kuu"
	"gopkg.in/guregu/null.v3"
)

type Bank struct {
	kuu.Model `rest:"*" displayName:"银行档案"`
	Name      null.String
}
