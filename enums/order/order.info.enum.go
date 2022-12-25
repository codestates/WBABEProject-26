package order_list_enums

/////////////////////////
//	Order Status
/////////////////////////

type OrderStatus int

const (
	OrderReceipt OrderStatus = iota +1  //주문 접수
	Cooking //조리중
	InDelivery //배달중
	DeliveryComplete //주문완료
	OrderAddChange //추가 주문으로 변경
	
)

/*
	Ordered //주문됨
	Cooked //조리완료
	OrderCancel //주문취소
	OrderReject //주문 거부
*/