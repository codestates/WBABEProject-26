package receipt_router

import (
	receipt_controller "wemade_project/controller/receipt"

	"github.com/gin-gonic/gin"
)

type MenuRoute struct {
	menuController receipt_controller.MenuController 
}

func InitWithSelf(menuController receipt_controller.MenuController ) MenuRoute {
	return MenuRoute{menuController: menuController}
}


func (r *MenuRoute) InitWithRoute(server *gin.Engine) {
	storeRouterV1 := server.Group("/api/v1/store")
	{
		storeRouterV1.GET("menu/get", r.menuController.GetMenu())
		storeRouterV1.POST("/menu/add",r.menuController.AddMenu())
		storeRouterV1.PUT("/menu/update", r.menuController.UpdateMenu())
	}
}



