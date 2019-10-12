package bk

import "github.com/kuuland/kuu"

func Mod() *kuu.Mod {
	return &kuu.Mod{
		Code: "bk",
		Models: []interface{}{
			&Book{},
			&Member{},
		},
		Routes: kuu.RoutesInfo{
			Public(),
			Private(),
			Signup(),
			Login(),
		},
	}
}
