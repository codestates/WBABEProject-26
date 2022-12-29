package user_route

import (
	user_controller "wemade_project/controller/user"

	"github.com/gin-gonic/gin"
)


type UserRoute struct {
	userController user_controller.UserController
}

func InitWithSelf(userController user_controller.UserController ) UserRoute {
	return UserRoute{userController: userController}
}

func (r *UserRoute) InitWithRoute(server *gin.Engine) {
	userRouterV1 := server.Group("/api/v1/account")
	{
		userRouterV1.POST("/user/add",r.userController.AddUser())
	}
}