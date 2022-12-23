package receipt_service

import (
	"time"

	"wemade_project/dto"
	receipt_model "wemade_project/model/receipt"

	"github.com/gofrs/uuid"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuService struct {
	memuCollection receipt_model.MenuCollection
}

func InitWithSelf(model receipt_model.MenuCollection) MenuService {
	return MenuService{memuCollection: model }
}



func (ms *MenuService) AddMenuItem(dto dto.CreateMenuRequest) (*dto.ReadMenuResponse, error) {
	//Model 조립
	var menuEntity receipt_model.MenuEntity
	menuEntity.CreateDate = time.Now()
	menuEntity.UpdateDate = menuEntity.CreateDate
	menuEntity.Id = uuid.Must(uuid.NewV4()).String()

	menuEntity.Name  = dto.Name
	menuEntity.IsCanOrder = dto.IsCanOrder
	menuEntity.Price = dto.Price
	menuEntity.Event = dto.Event
	menuEntity.MenuCategory = dto.MenuCategory

	var subMenu []receipt_model.SubMenu
	for _, val := range dto.SubMenu {
		item := receipt_model.SubMenu{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	menuEntity.SubMenu = subMenu
	menuEntity.FoodEtcInfo = receipt_model.FoodEtcInfo(dto.FoodEtcInfo)

	//데이터 등록 처리
	result, err := ms.memuCollection.AddEntity(menuEntity);
	if err != nil {
		return nil, err
	}

	//등록된 아이템을 반환하기 위해 조회
	saveItem, err1 := ms.memuCollection.FindByInnerId(result.InsertedID); 
	if err1 != nil {
		return nil, err1
	}

	return changeMenuEntity2ReadDto(*saveItem), nil
}


func changeMenuEntity2ReadDto(entity receipt_model.MenuEntity) *dto.ReadMenuResponse {
	var subMenu []dto.CreateSubMenuRequest
	for _, val := range entity.SubMenu {
		item := dto.CreateSubMenuRequest{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	return &dto.ReadMenuResponse{Id: entity.Id, Name:  entity.Name, IsCanOrder: entity.IsCanOrder, Price: entity.Price, Event: entity.Event, MenuCategory:  entity.MenuCategory, SubMenu:  subMenu, FoodEtcInfo: dto.CreateFoodEtcInfoRequest(entity.FoodEtcInfo) }
}

/*
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
*/