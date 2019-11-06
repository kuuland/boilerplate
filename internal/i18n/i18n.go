package i18n

import (
	"github.com/kuuland/kuu"
)

// InitLanguageMessages
func InitLanguageMessages() {
	register := kuu.NewLangRegister(kuu.DB())

	// Bank
	register.SetKey("bank_name").Add("Bank Name", "银行名称", "銀行名稱")
	// BankCard
	register.SetKey("bank_card_bank").Add("Bank", "所属银行", "所屬銀行")
	register.SetKey("bank_card_customer").Add("Customer", "所属客户", "所屬客戶")
	register.SetKey("bank_card_cardsn").Add("Card No.", "卡号", "卡號")
	_ = register.Exec(true)
}
