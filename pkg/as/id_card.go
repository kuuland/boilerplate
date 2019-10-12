package as

import (
	"github.com/kuuland/kuu"
	"gopkg.in/guregu/null.v3"
)

type IDCard struct {
	kuu.Model  `rest:"*" displayName:"身份证"`
	CustomerID uint
	IDNum      null.String
}
