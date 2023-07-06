package routers

import (
	controllers "test_jenkins/backend/httpserver/controllers"
	"test_jenkins/backend/httpserver/middleware"
	"test_jenkins/backend/utils"

	"github.com/gin-gonic/gin"
)

func UserRouter(route *gin.RouterGroup, userController controllers.UserController, authService utils.AuthHelper) *gin.RouterGroup {
	userRouter := route.Group("/users")
	{
		userRouter.POST("register", userController.Register)
		userRouter.POST("login", userController.Login)
		userRouter.GET("", middleware.JwtGuard(authService), userController.GetUsers)
		userRouter.PUT("update-account", middleware.JwtGuard(authService), userController.UpdateUser)
		userRouter.DELETE("delete-account", middleware.JwtGuard(authService), userController.DeleteUser)
	}
	return userRouter
}
