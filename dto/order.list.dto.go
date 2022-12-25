package dto

import (
	order_enums "wemade_project/enums/order"
	menu_model "wemade_project/model/menu"
)

/////////////////////////
//	 Create Request
/////////////////////////

type CreateOrderListRequest struct {
	OrderUserId string `json:"orderUserId" binding:"required"`
	OrderMenuList []string `json:"orderMenuList" binding:"required"` //주문 메뉴 리스트
}


/////////////////////////
//		Resposne
/////////////////////////


//Normal version Order List
type NomalReadOrderListResponse struct {
	OrderId string `json:"orderId"` //고유 id
	OrderUserId string `json:"orderUserId"` //주문자
	OrderMenu []menu_model.MenuEntity `json:"orderMenu"` //주문 메뉴 리스트
	OrderStatus order_enums.OrderStatus `json:"orderStatus"` //주문 상태
	TotalPrice int `json:"totalPrice"` //총 가격
}

/*
`json:"orderUserId", binding:"required"` //주문자
type CreateUserRequest struct {
	Name string `json:"name" binding:"required" ` //사용자 이름 
	Phone string `json:"phone" binding:"required"` //사용자 폰번호
	Addr string `json:"addr" binding:"required"` //주소
}
type OrderListEntity struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	OrderId string `bson:"orderId"` //고유 id
	OrderUser string `bson:"orderUserId"` //주문자
	OrderMenu []menu_model.MenuEntity `bson:"orderMenu"` //주문 메뉴 리스트
	OrderStatus order_enums.OrderStatus `bosn:"orderStatus"` //주문 상태
	TotalPrice int `bson:"totalPrice"` //총 가격
	CreateDate time.Time `bson:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `bson:"updateDate"` //데이터 수정 시각
}
*/