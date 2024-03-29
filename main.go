package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kuuland/boilerplate/middlewares"
	"github.com/kuuland/boilerplate/pkg"
	"github.com/kuuland/kuu"
	"regexp"
)

func main() {
	r := kuu.Default()
	r.RegisterWhitelist(
		"GET /book/public",
		"POST /member/login",
		regexp.MustCompile(`book$`),
		regexp.MustCompile(`bank`),
		regexp.MustCompile(`customer$`),
		regexp.MustCompile(`idcard$`),
	)
	r.Use(middlewares.HelloMiddleware)
	r.Import(kuu.Acc(), kuu.Sys(), pkg.All())
	r.Run()
}
