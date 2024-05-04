package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	BookRouter
	UserRouter
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("api")
	var router = new(Router)

	router.InitBookRouter(group)
	router.InitUserRouter(group)

	return r
}
