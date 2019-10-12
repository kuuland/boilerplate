package bk

import (
	"github.com/jinzhu/gorm"
	"github.com/kuuland/kuu"
)

// Member
type Member struct {
	kuu.Model `rest:"*" displayName:"会员档案"`
	Mobile    string `name:"手机号"`
	Email     string `name:"邮箱地址"`
	Password  string `name:"密码" kuu:"password" binding:"required"`
	UID       uint
}

// BeforeSave
func (m *Member) BeforeSave(scope *gorm.Scope) (err error) {
	if len(m.Password) == 32 {
		var hashed string
		if hashed, err = kuu.GenerateFromPassword(m.Password); err == nil {
			err = scope.SetColumn("Password", hashed)
		}
	}
	return
}
