package menu_router

import (
	menu_controller "wemade_project/controller/menu"

	"github.com/gin-gonic/gin"
)

type MenuRoute struct {
	menuController menu_controller.MenuController 
}

func InitWithSelf(menuController menu_controller.MenuController ) MenuRoute {
	return MenuRoute{menuController: menuController}
}


func (r *MenuRoute) InitWithRoute(server *gin.Engine) {
	storeMenuRouterV1 := server.Group("/api/v1/store")
	{
		storeMenuRouterV1.GET("menu/get", r.menuController.GetMenuList())
		storeMenuRouterV1.GET("menu/get/:menu_id", r.menuController.GetMenu4MenuId())
		storeMenuRouterV1.POST("/menu/add",r.menuController.AddMenu())
		storeMenuRouterV1.PUT("/menu/update", r.menuController.UpdateMenu())
		storeMenuRouterV1.DELETE("/menu/delete/:menu_id", r.menuController.DeleteMenu4Logical())
	}
}



