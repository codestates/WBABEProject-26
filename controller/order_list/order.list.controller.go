package order_list_controller

import order_list_service "wemade_project/service/order_list"

/////////////////////////
//	  Struct
/////////////////////////

type OrderListController struct {
	orderListService order_list_service.OrderListService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(orderListService order_list_service.OrderListService) OrderListController {
	return OrderListController{orderListService: orderListService}
}
