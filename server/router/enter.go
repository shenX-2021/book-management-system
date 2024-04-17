package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	BookRouter
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	group := r.Group("api")
	var router = new(Router)

	router.InitBookRouter(group)

	return r
}
