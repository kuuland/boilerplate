package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kuuland/boilerplate/internal/book"
	"github.com/kuuland/boilerplate/internal/member"
	"github.com/kuuland/boilerplate/middlewares"
	"github.com/kuuland/kuu"
	"regexp"
)

func main() {
	r := kuu.Default()
	r.RegisterWhitelist(
		"GET /book/public",
		"POST /member/login",
		regexp.MustCompile(`book$`),
	)
	r.Import(kuu.Acc(), kuu.Sys(), &kuu.Mod{
		Code: "foo",
		Models: []interface{}{
			&book.Book{},
			&member.Member{},
		},
		Middlewares: gin.HandlersChain{
			middlewares.HelloMiddleware,
		},
		Routes: kuu.RoutesInfo{
			book.Private(),
			book.Public(),
			member.Signup(),
			member.Login(),
		},
	})
	r.Run()
}
