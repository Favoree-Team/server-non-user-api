package routes

import (
	"github.com/Favoree-Team/server-non-user-api/controller"
	"github.com/gin-gonic/gin"
)

func PingRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.Ping)
	}
}
