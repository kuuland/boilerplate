package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kuuland/boilerplate/internal/book"
	"github.com/kuuland/boilerplate/middlewares"
	"github.com/kuuland/boilerplate/models"
	"github.com/kuuland/boilerplate/routes"
	"github.com/kuuland/kuu"
	"regexp"
)

func main() {
	r := kuu.Default()
	r.RegisterWhitelist(
		"GET /book/public",
		regexp.MustCompile(`book$`),
	)
	r.Import(kuu.Acc(), kuu.Sys(), &kuu.Mod{
		Code: "foo",
		Models: []interface{}{
			&book.Book{},
		},
		Middlewares: gin.HandlersChain{
			middlewares.HelloMiddleware,
		},
		Routes: kuu.RoutesInfo{
			book.BookPrivate(),
			book.BookPublic(),
		},
	})
	r.Run()
}
