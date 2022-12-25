package dto

import (
	order_enums "wemade_project/enums/order"
)

/////////////////////////
//	 Create Request
/////////////////////////

/**
* 초기 주문 요청 리퀘스트
 */
type CreateOrderListRequest struct {
	OrderUserId string `json:"orderUserId" binding:"required"`
	OrderMenuList []string `json:"orderMenuList" binding:"required"` //주문 메뉴 리스트
}


/**
* 주문 요청 업데이트 리퀘스트
*/
type UpdateOrderListRequest struct {
	OrderId string `json:"orderId" binding:"required"` //고유 id
	OrderMenu []string `json:"orderMenu" binding:"required"` //주문 메뉴 리스트
}


/////////////////////////
//		Resposne
/////////////////////////


//Normal version Order List
type NomalReadOrderListResponse struct {
	OrderId string `json:"orderId"` //고유 id
	OrderUserId string `json:"orderUserId"` //주문자
	OrderMenu []NormalReadMenuResponse `json:"orderMenu"` //주문 메뉴 리스트
	OrderStatus order_enums.OrderStatus `json:"orderStatus"` //주문 상태
	TotalPrice int `json:"totalPrice"` //총 가격
}

