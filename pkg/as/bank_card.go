package as

import (
	"github.com/kuuland/kuu"
	"gopkg.in/guregu/null.v3"
)

type BankCard struct {
	kuu.Model  `rest:"*" displayName:"银行卡"`
	PubBank    Bank        `name:"所属银行" gorm:"foreignkey:BankID"`
	BankID     uint        `name:"所属银行ID"`
	CustomerID uint        `name:"所属客户ID"`
	CardNum    null.String `name:"卡号"`
}
