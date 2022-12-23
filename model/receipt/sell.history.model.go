package receipt_model

import (
	"time"
	order_enums "wemade_project/enums/order"
)

type SellHistory struct {
	OrderId string //주문 고유 id
	StoreId string //판매 사업장 id
	OdererId string //주문자 id
	TotalPrice int //총 판매금
	OrderStatus order_enums.OrderStatus //주문상태
	OrderMenu []MenuEntity //주문 메뉴
	CreateDate time.Time //데이터 생성 시각
	UpdateDate time.Time //데이터 수정 시각
}
