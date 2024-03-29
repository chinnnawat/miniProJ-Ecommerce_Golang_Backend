package routes

import "github.com/gin-gonic/gin"

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("/user/signup")
	incommingRoutes.POST("/user/login")
	incommingRoutes.POST("/admin/addproduct")
	incommingRoutes.GET("/users/productview")
	incommingRoutes.GET("/users/search")
}
