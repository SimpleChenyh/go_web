package router

import (
	"chenyh.com/go_web/routerHandle"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine) {

	userR := r.Group("user/")

	userR.GET(":id", routerHandle.GetUser)


}
