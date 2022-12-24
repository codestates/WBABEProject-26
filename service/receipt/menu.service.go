package receipt_service

import (
	"time"

	"wemade_project/dto"
	receipt_model "wemade_project/model/receipt"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuService struct {
	memuCollection receipt_model.MenuCollection
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model receipt_model.MenuCollection) MenuService {
	return MenuService{memuCollection: model }
}


/////////////////////////
//	  Add Data
/////////////////////////

func (ms *MenuService) AddMenuItem(addDto dto.CreateMenuRequest) (*dto.ReadMenuResponse, error) {
	//Model 조립
	// var menuEntity receipt_model.MenuEntity
	menuEntity := changeMenuRequest2Entity(dto.MenuRequest(addDto))
	menuEntity.CreateDate = time.Now()
	menuEntity.UpdateDate = menuEntity.CreateDate
	menuEntity.Id = uuid.Must(uuid.NewV4()).String()

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


/////////////////////////
//	  Update
/////////////////////////

//메뉴 업데이트
func (ms *MenuService) UpdateMenuItem(sendDto dto.UpdateMenuRequest) (*dto.ReadMenuResponse, error) {
	//Model 조립
	// var menuEntity receipt_model.MenuEntity
	
	//Id를 통해 수정 대상 메뉴 아이템을 찾아온다.
	findMenu, findErr := ms.memuCollection.FindByMenuId(sendDto.Id)
	if findErr != nil {
		return nil, findErr
	}
	
	//수정할 데이터 셋
	var setUpdateSet bson.D

	if (sendDto.Name != "") {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "name", Value: sendDto.Name})
	}
	if sendDto.IsCanOrder != 0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "isCanOrder", Value: sendDto.IsCanOrder})
	}
	if sendDto.Price != 0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "price", Value: sendDto.Price})
	}	
	if len(sendDto.Event) !=0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "event", Value: sendDto.Event})
	}
	if len(sendDto.MenuCategory) !=0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "menuCategory", Value: sendDto.MenuCategory})
	}
	if len(sendDto.SubMenu) !=0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "subMenu", Value: sendDto.SubMenu})
	}
	if (sendDto.FoodEtcInfo != dto.FoodEtcInfoRequest{}) {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "foodEtcInfo", Value: sendDto.FoodEtcInfo})
	}

	//수정시간을 변경한다.
	setUpdateSet = append(setUpdateSet, bson.E{Key: "updateDate", Value: time.Now()})

		
	

	saveItem, err1 := ms.memuCollection.UpdateEntity(findMenu.ID, setUpdateSet)
	if err1 != nil {
		return nil, err1
	}

	return changeMenuEntity2ReadDto(*saveItem), nil
} 


/////////////////////////
//	  Utils
/////////////////////////

//메뉴 엔티티를 읽기 DTO로 변환시키는 함수
func changeMenuEntity2ReadDto(entity receipt_model.MenuEntity) *dto.ReadMenuResponse {
	var subMenu []dto.SubMenuRequest
	for _, val := range entity.SubMenu {
		item := dto.SubMenuRequest{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	return &dto.ReadMenuResponse{Id: entity.Id, Name:  entity.Name, IsCanOrder: entity.IsCanOrder, Price: entity.Price, Event: entity.Event, MenuCategory:  entity.MenuCategory, SubMenu:  subMenu, FoodEtcInfo: dto.FoodEtcInfoRequest(entity.FoodEtcInfo) }
}

//공통 MenuRequest 데이터 중 일부를 Entity랑 매핑하는 함수
func changeMenuRequest2Entity(sendDto dto.MenuRequest) receipt_model.MenuEntity {
	var menuEntity receipt_model.MenuEntity
	
	menuEntity.Name  = sendDto.Name
	menuEntity.IsCanOrder = sendDto.IsCanOrder
	menuEntity.Price = sendDto.Price
	menuEntity.Event = sendDto.Event
	menuEntity.MenuCategory = sendDto.MenuCategory

	// var subMenu []receipt_model.SubMenu
	// for _, val := range sendDto.SubMenu {
	// 	item := receipt_model.SubMenu{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
	// 	subMenu = append(subMenu, item)
	// }

	menuEntity.SubMenu = changeSubMenu4Dto2Entity(sendDto.SubMenu)
	menuEntity.FoodEtcInfo = receipt_model.FoodEtcInfo(sendDto.FoodEtcInfo)

	return menuEntity;
}


//Menu 안의 SubMenu 의 dto형태를 entity 형태로 맞춰주는 함수
func changeSubMenu4Dto2Entity(sendDto []dto.SubMenuRequest)[]receipt_model.SubMenu {
	var subMenu []receipt_model.SubMenu
	for _, val := range sendDto {
		item := receipt_model.SubMenu{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}
	return subMenu
}

