package book

import "github.com/kuuland/kuu"

// Mod
func Mod() *kuu.Mod {
	return &kuu.Mod{
		Code: "book",
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
