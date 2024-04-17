package router

import (
	"github.com/gin-gonic/gin"

	controller "book-management-system-server/controller"
)

type BookRouter struct{}

func (s *BookRouter) InitBookRouter(g *gin.RouterGroup) {
	var ctrl = new(controller.BookCtrl)
	g.POST("book/create", ctrl.Create)
	g.GET("book/list", ctrl.List)
	g.PATCH("book/:id", ctrl.Edit)
	g.DELETE("book/:id", ctrl.Delete)
}
