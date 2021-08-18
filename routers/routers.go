package routers

import (
	"github.com/gin-gonic/gin"
	"tma.com.vn/lngochuy/rest/controllers"
)

func Route() *gin.Engine {
	router := gin.Default()

	router.GET("users", controllers.GetUserList)
	router.POST("users", controllers.CreateUser)
	router.GET("users/:id", controllers.GetUser)
	router.PUT("users/:id", controllers.UpdateUser)
	router.DELETE("users/:id", controllers.RemoveUser)

	return router
}
