package member

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/kuuland/kuu"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Signup
func Signup() kuu.RouteInfo {
	return kuu.RouteInfo{
		Path:   "/member/signup",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var memberDoc Member
			if err := c.ShouldBindJSON(&memberDoc); err != nil {
				c.STDErr("解析请求体失败，请按照Member格式提交参数", err)
				return
			}
			err := c.WithTransaction(func(tx *gorm.DB) (err error) {
				// 创建会员档案
				// emberDoc.Password = kuu.MD5(models.GenerateRandomPass())
				err = tx.Create(&memberDoc).Error
				if err != nil {
					return
				}
				if memberDoc.ID == 0 {
					return errors.New("创建会员失败")
				}
				// 创建并绑定预置用户档案
				memberUser := kuu.User{
					Username: uuid.NewV4().String(),
					Password: memberDoc.Password,
					SubDocID: memberDoc.ID,
					Mobile:   memberDoc.Mobile,
					Email:    memberDoc.Email,
				}
				err = tx.Create(&memberUser).Error
				if err != nil {
					return
				}
				if memberUser.ID == 0 {
					return errors.New("创建会员帐号失败")
				}
				tx.Model(&memberDoc).Where("id = ? ", memberDoc.ID).Update("uid", memberUser.ID)
				return tx.Error
			})
			if err != nil {
				c.STDErr("注册失败", err)
			} else {
				c.STD(memberDoc)
			}
		},
	}
}

// Login
func Login() kuu.RouteInfo {
	return kuu.RouteInfo{
		Path:   "/member/login",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var login struct {
				User string
				Pass string
			}
			if err := c.ShouldBindJSON(&login); err != nil {
				c.STDErr("解析请求体失败，请按照Login格式提交参数", err)
				return
			}
			var member Member
			kuu.DB().Where("mobile = ? or email = ?", login.User, login.User).Find(&member)
			if kuu.CompareHashAndPassword(member.Password, login.Pass) {
				secret, err := kuu.GenToken(kuu.GenTokenDesc{
					Payload: jwt.MapClaims{
						"MemberID":  member.ID,
						"CreatedAt": member.CreatedAt,
					},
					UID:        member.UID,
					SubDocID:   member.ID,
					Expiration: time.Second * time.Duration(kuu.ExpiresSeconds),
				})
				if err != nil {
					c.STDErr("登录失败: 无法生成token")
					return
				}
				c.STD(secret.Token)
				return
			} else {
				c.STD("登陆失败用户名或者密码错误!")
			}
		},
	}
}
