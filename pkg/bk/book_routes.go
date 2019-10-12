package bk

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/kuuland/kuu"
	"strconv"
)

// Private
func Private() kuu.RouteInfo {
	return kuu.RouteInfo{
		Name:   "查询私有图书",
		Path:   "/book/private",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var (
				body struct {
					Subject string `json:"Name" binding:"required"`
					InStock bool
				}
				books         []Book
				failedMessage = c.L("book_private_failed", "Querying books failed")
			)

			if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
				c.STDErr(failedMessage, err)
				return
			}

			if err := c.DB().Where(body).Find(&books).Error; err != nil {
				c.STDErr(failedMessage, err)
				return
			}
			c.STD(books)
		},
	}
}

// Public
func Public() kuu.RouteInfo {
	return kuu.RouteInfo{
		Name:   "查询公共图书",
		Path:   "/book/public",
		Method: "GET",
		HandlerFunc: func(c *kuu.Context) {
			var (
				books         []Book
				class, _      = strconv.Atoi(c.DefaultQuery("c", "100"))
				failedMessage = c.L("book_public_failed", "Querying books failed")
			)

			if err := c.DB().Where(Book{InStock: true, Class: class}).Find(&books).Error; err != nil {
				c.STDErr(failedMessage, err)
				return
			}
			c.STD(books)
		},
	}
}
