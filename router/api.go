package router

import (
	"web/controller/HelloController"
	"web/controller/LoginController"

	_ "web/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoute(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/hello", HelloController.Hello)
	v1 := r.Group("/user")
	{
		v1.POST("/register", LoginController.Register)
	}
}
