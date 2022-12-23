package receipt_model

import (
	enum_order "wemade_project/enums/order"
)

type OperatingTime struct {
	OpenTime string `` //오픈시간
	CloseTime string `` //종료시간
	BreakTime string `json:"breakTime,omitempty"` //브레이크 타임
}

//사업장 정보
type Store struct {
	Name	string //상호명
	Address	string //주소
	Contact	string //연락처
	RunTime	OperatingTime //운영시간
	Menus []MenuEntity //판매 메뉴
	StoreType []enum_order.StoreCategory //사업장 업종 타입 (복수)
	SellHistory []SellHistory //주문 히스토리
}
