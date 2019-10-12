package bk

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/kuuland/kuu"
	uuid "github.com/satori/go.uuid"
	"time"
)

// Signup
func Signup() kuu.RouteInfo {
	return kuu.RouteInfo{
		Name:   "会员注册接口",
		Path:   "/member/signup",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var memberDoc Member
			failedMessage := c.L("member_signup_failed", "Member signup failed")
			if err := c.ShouldBindJSON(&memberDoc); err != nil {
				c.STDErr(failedMessage, err)
				return
			}
			err := c.WithTransaction(func(tx *gorm.DB) error {
				// 创建会员档案
				// emberDoc.Password = kuu.MD5(models.GenerateRandomPass())
				if err := tx.Create(&memberDoc).Error; err != nil {
					return err
				}
				// 创建并绑定预置用户档案
				memberUser := kuu.User{
					Username: uuid.NewV4().String(),
					Password: memberDoc.Password,
					SubDocID: memberDoc.ID,
					Mobile:   memberDoc.Mobile,
					Email:    memberDoc.Email,
				}

				if err := tx.Create(&memberUser).Error; err != nil {
					return err
				}
				tx.Model(&memberDoc).Where("id = ? ", memberDoc.ID).Update("uid", memberUser.ID)
				return tx.Error
			})
			if err != nil {
				c.STDErr(failedMessage, err)
			} else {
				c.STD(memberDoc)
			}
		},
	}
}

// Login
func Login() kuu.RouteInfo {
	return kuu.RouteInfo{
		Name:   "会员登录接口",
		Path:   "/member/login",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var login struct {
				User string
				Pass string
			}
			failedMessage := c.L("member_login_failed", "Member login failed")
			if err := c.ShouldBindJSON(&login); err != nil {
				c.STDErr(failedMessage, err)
				return
			}
			var member Member
			kuu.DB().Where("mobile = ? or email = ?", login.User, login.User).Find(&member)
			if err := kuu.CompareHashAndPassword(member.Password, login.Pass); err != nil {
				c.STDErr(failedMessage, err)
				return
			} else {
				secret, err := kuu.GenToken(kuu.GenTokenDesc{
					Payload: jwt.MapClaims{
						"MemberID":  member.ID,
						"CreatedAt": member.CreatedAt,
					},
					UID:      member.UID,
					SubDocID: member.ID,
					Exp:      time.Now().Add(time.Second * time.Duration(kuu.ExpiresSeconds)).Unix(),
					Type:     "MY_SIGN_TYPE",
				})
				if err != nil {
					c.STDErr(failedMessage, err)
					return
				}
				c.STD(secret.Token)
			}
		},
	}
}
