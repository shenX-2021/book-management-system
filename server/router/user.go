package router

import (
	"github.com/gin-gonic/gin"

	controller "book-management-system-server/controller"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(g *gin.RouterGroup) {
	var ctrl = new(controller.UserCtrl)
	g.POST("user/register", ctrl.Register)
	g.POST("user/login", ctrl.Login)
}
