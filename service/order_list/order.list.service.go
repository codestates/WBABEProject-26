package order_list_service

import (
	order_list_model "wemade_project/model/order"
)

/////////////////////////
//	  Struct
/////////////////////////

type OrderListService struct {
	orderListCollection order_list_model.OrderListCollection
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model order_list_model.OrderListCollection) OrderListService {
	return OrderListService{orderListCollection: model }
}