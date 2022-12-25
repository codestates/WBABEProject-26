package order_list_router

import (
	order_list_controller "wemade_project/controller/order_list"

	"github.com/gin-gonic/gin"
)


type OrderListRoute struct {
	orderListController order_list_controller.OrderListController
}

func InitWithSelf(orderListController order_list_controller.OrderListController ) OrderListRoute {
	return OrderListRoute{orderListController: orderListController}
}

func (r *OrderListRoute) InitWithRoute(server *gin.Engine) {
	orderListRouterV1 := server.Group("/api/v1/order_list")
	{
		//
		orderListRouterV1.GET("order/user/:user_id", r.orderListController.Find4OrderUserId())
		orderListRouterV1.GET("order", r.orderListController.Find4All())
		
		//주문 접수
		orderListRouterV1.POST("/order/add", r.orderListController.AddOrderListItem())
		orderListRouterV1.PUT("/order/menu/update", r.orderListController.UpdateOrderList4Menu())
		orderListRouterV1.PUT("/order/status/update", r.orderListController.UpdateOrderList4Status())

		
		// orderListRouterV1.PUT("/menu/update", r.menuController.UpdateMenu())
		// orderListRouterV1.DELETE("/menu/delete/:menu_id", r.menuController.DeleteMenu4Logical())
	}
}