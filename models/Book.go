package models

import (
	"github.com/kuuland/kuu"
	"time"
)

// 枚举：图书分类
const (
	_                   = iota
	BookClassPhilosophy = 100 * iota // 100的倍数
	BookClassSocial
	BookClassNatural
)

func init() {
	kuu.Enum("BookClass", "图书分类").
		Add(BookClassPhilosophy, "哲学").
		Add(BookClassSocial, "社会科学").
		Add(BookClassNatural, "自然科学")
}

// Book
type Book struct {
	kuu.Model `rest:"*" displayName:"图书信息"`
	Subject   string     `name:"书名"`
	Intro     string     `name:"简介"`
	Class     int        `name:"分类" enum:"BookClass"`
	Price     float32    `name:"价格"`
	InStock   bool       `name:"是否有货"`
	PubDate   *time.Time `name:"出版日期"`
}
