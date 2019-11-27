package core

import (
	"github.com/kuuland/kuu"
	"gopkg.in/guregu/null.v3"
)

type Customer struct {
	kuu.Model `rest:"*" displayName:"客户档案"`
	Name      null.String `name:"客户姓名"`
	RefIDCard IDCard      `name:"拥有身份证"`
	BankCards []BankCard  `name:"拥有银行卡"`
	RefBanks  []Bank      `name:"关联银行" gorm:"many2many:as_Customer2Banks"`
}
