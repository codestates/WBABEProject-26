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

/*
엔드포인트 구성에 대해서 전반적인 코멘트 드립니다.
1. REST API의 성숙도 모델에 대해서 공부해보시면 좋을 것 같습니다.

2. 일반적으로 HTTP URI에 new, modify 와 같은 행위는 들어가지 않습니다. 
	복수형의 단어로 구성을 하고, 동일한 URI 내에서 http method만 변경하여 행위를 표현하는 것이 일반적인 REST API의 구성 방식입니다.
	e.g.
	GET v1/menus -> 메뉴 목록을 조회.
	GET v1/menus/1 -> 1번 메뉴를 조회.
	POST v1/menus -> 메뉴를 생성.
	PATCH v1/menus/1 -> 1번 메뉴에 대해서 업데이트
	DELETE v1/menus/1 -> 1번 메뉴에 대해서 삭제
*/

/*
각 라우팅을 주제에 맞게 다른 파일로 분리해주신 점 좋습니다.
이렇게 구성하면 추후 확장성에 있어서도 용이해 보입니다.
*/

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



