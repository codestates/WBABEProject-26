package order_enums

type OrderStatus int

const (
	Ordered OrderStatus = iota +1 //주문됨
	OrderReceipt //주문 접수
	OrderReject //주문 거부
	Cooking //조리중
	Cooked //조리완료
	InDelivery //배달중
	DeliveryComplete //주문완료
)