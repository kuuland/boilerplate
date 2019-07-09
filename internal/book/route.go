package book

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/kuuland/boilerplate/models"
	"github.com/kuuland/kuu"
	"strconv"
)

// BookPrivate
func BookPrivate() kuu.RouteInfo {
	return kuu.RouteInfo{
		Path:   "/book/private",
		Method: "POST",
		HandlerFunc: func(c *kuu.Context) {
			var body struct {
				Subject string `json:"Name" binding:"required"`
				InStock bool
			}
			if err := c.ShouldBindBodyWith(&body, binding.JSON); err != nil {
				c.STDErr("解析请求体失败", err)
				return
			}
			var books []models.Book
			if err := c.DB().Find(&books).Error; err != nil {
				c.STDErr("查询失败", err)
				return
			}
			c.STD(books)
		},
	}
}

// BookPublic
func BookPublic() kuu.RouteInfo {
	return kuu.RouteInfo{
		Path:   "/book/public",
		Method: "GET",
		HandlerFunc: func(c *kuu.Context) {
			var books []models.Book
			class, _ := strconv.Atoi(c.DefaultQuery("c", "100"))
			if err := c.DB().Where(models.Book{InStock: true, Class: class}).Find(&books).Error; err != nil {
				c.STDErr("查询失败", err)
				return
			}
			c.STD(books)
		},
	}
}
