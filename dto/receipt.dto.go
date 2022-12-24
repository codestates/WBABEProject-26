package dto

import (
	"time"
	receipt_enums "wemade_project/enums/receipt"
)

/////////////////////////
// Request & Response
/////////////////////////

//음식 기타 정보
type FoodEtcInfoRequest struct {
	OriginInfo string `json:"originInfo,omitempty"` //원산지
	SpicyInfo receipt_enums.FoodSpicyType `json:"spicyInfo,omitempty" binding:"omitempty"` //맵기 정보  binding:"required"
}

//서브 메뉴
type SubMenuRequest struct {
	SubMenuName string `json:"subMenuName,omitempty"` //서브 메뉴 타이틀
	Name string `json:"name,omitempty"` //서브 메뉴 이름
	Price int `json:"price,omitempty"` //가격
}

//메뉴 공통
type MenuRequest struct {
	Name string `json:"name,omitempty" binding:"checkerMenuEvent"` //메뉴 이름
	MenuStatus receipt_enums.MenuSellStatusType `json:"menuStatus,omitempty"`	//주문 가능 여부	
	Price int `json:"price,omitempty"` //가격
	Event []receipt_enums.MenuEventType `json:"event,omitempty" binding:"checkerMenuEvent"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory,omitempty" ` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu,omitempty"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"etcInfo,omitempty"`
}

/*

*/

/////////////////////////
//	 Create Request
/////////////////////////

//메뉴 생성에 사용하는 데이터
type CreateMenuRequest struct {
	Name string `json:"name" binding:"required" ` //메뉴 이름
	MenuStatus receipt_enums.MenuSellStatusType `json:"menuStatus" binding:"required"`	//주문 가능 여부	
	Price int `json:"price" binding:"required"` //가격
	Event []receipt_enums.MenuEventType `json:"event" binding:"required,checkerMenuEvent"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory" binding:"required"` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu" binding:"required"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"etcInfo" binding:"required"`
}


/////////////////////////
//	 Update Request
/////////////////////////

//수정할 때 사용하는 dto
type UpdateMenuRequest struct {
	Id string `json:"id" binding:"required"` //해당 메뉴 고유 id
	Name string `json:"name" binding:"-"` //메뉴 이름
	MenuStatus receipt_enums.MenuSellStatusType `json:"menuStatus" binding:"-"`	//주문 가능 여부	
	Price int `json:"price" binding:"-"` //가격
	Event []receipt_enums.MenuEventType `json:"event" binding:"checkerMenuEvent"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory" ` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu" binding:"omitempty"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"etcInfo" binding:"omitempty"`
}

/*
Id string `json:"id" binding:"required"` //해당 메뉴 고유 id
	Name string `json:"name,omitempty" binding:"checkerMenuEvent"` //메뉴 이름
	IsCanOrder receipt_enums.MenuSellStatusType `json:"isCanOrder,omitempty"`	//주문 가능 여부	
	Price int `json:"price,omitempty"` //가격
	Event []receipt_enums.MenuEventType `json:"event,omitempty" binding:"checkerMenuEvent"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory,omitempty" ` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu,omitempty" binding:"omitempty"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"etcInfo,omitempty" binding:"omitempty"`
*/

/////////////////////////
//		Resposne
/////////////////////////

//피주문자 메뉴 열람용 리스폰스
type ReceiptReadMenuResponse struct {
	Id string `json:"id"`
	Name string `json:"name"` //메뉴 이름
	MenuStatus receipt_enums.MenuSellStatusType `json:"menuStaus"`	//주문 가능 여부	
	Price int `json:"price"` //가격
	Event []receipt_enums.MenuEventType `json:"event"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory"` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"foodEtcInfo"` //기타 정보
	CreateDate time.Time `json:"createDate"` //데이터 생성 시각
	UpdateDate time.Time `json:"updateDate"` //데이터 수정 시각
}

type UserReadMenuResponse struct {
	Id string `json:"id"`
	Name string `json:"name"` //메뉴 이름
	MenuStatus receipt_enums.MenuSellStatusType `json:"menuStaus"`	//주문 가능 여부	
	Price int `json:"price"` //가격
	Event []receipt_enums.MenuEventType `json:"event"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory"` //매뉴 카테고리
	SubMenu []SubMenuRequest `json:"subMenu"` //서브메뉴
	FoodEtcInfo FoodEtcInfoRequest `json:"foodEtcInfo"` //기타 정보
}
