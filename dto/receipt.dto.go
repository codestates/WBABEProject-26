package dto

import (
	receipt_enums "wemade_project/enums/receipt"
)

//음식 기타 정보
type CreateFoodEtcInfoRequest struct {
	OriginInfo string `json:"originInfo" binding:"required"` //원산지
	SpicyInfo receipt_enums.FoodSpicyType `json:"spicyInfo"` //맵기 정보  binding:"required"
}

type CreateSubMenuRequest struct {
	SubMenuName string `json:"subMenuName" binding:"required"` //서브 메뉴 타이틀
	Name string `json:"name" binding:"required"` //서브 메뉴 이름
	Price int `json:"price" binding:"required"` //가격
}

//메뉴 생성에 사용하는 데이터
type CreateMenuRequest struct {
	Name string `json:"name" binding:"required" ` //메뉴 이름
	IsCanOrder bool `json:"isCanOrder" `	//주문 가능 여부	
	Price int `json:"price" binding:"required"` //가격
	Event []receipt_enums.MenuEventType `json:"event" binding:"required,checkerMenuEvent"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory" binding:"required"` //매뉴 카테고리
	SubMenu []CreateSubMenuRequest `json:"subMenu" binding:"required"` //서브메뉴
	FoodEtcInfo CreateFoodEtcInfoRequest `json:"etcInfo" binding:"required"`
}


/////////////////////////
//		Resposne
/////////////////////////

type ReadMenuResponse struct {
	Id string `json:"id"`
	Name string `json:"name"` //메뉴 이름
	IsCanOrder bool `json:"isCanOrder"`	//주문 가능 여부	
	Price int `json:"price"` //가격
	Event []receipt_enums.MenuEventType `json:"event"` //이벤트
	MenuCategory []receipt_enums.MenuCategoryType `json:"menuCategory"` //매뉴 카테고리
	SubMenu []CreateSubMenuRequest `json:"subMenu"` //서브메뉴
	FoodEtcInfo CreateFoodEtcInfoRequest `json:"foodEtcInfo"` //기타 정보
}
