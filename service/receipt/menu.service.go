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
//	  Find Data
/////////////////////////

//메뉴 id로 엔티티 데이터를 조회하는 함수
func (ms *MenuService) Find4MenuId(id string) (*dto.ReceiptReadMenuResponse, error) {
	findMenu, findErr := ms.memuCollection.FindByMenuId(id)
	if findErr != nil {
		return nil, findErr
	}
	
	return changeMenuEntity2ReceiptReadDto(*findMenu), nil
}

/////////////////////////
//	  Add Data
/////////////////////////

func (ms *MenuService) AddMenuItem(addDto dto.CreateMenuRequest) (*dto.ReceiptReadMenuResponse, error) {
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
	saveItem, err1 := ms.memuCollection.FindByObjectId(result.InsertedID); 
	if err1 != nil {
		return nil, err1
	}

	return changeMenuEntity2ReceiptReadDto(*saveItem), nil
}


/////////////////////////
//	  Update
/////////////////////////

//메뉴 업데이트
func (ms *MenuService) UpdateMenuItem(sendDto dto.UpdateMenuRequest) (*dto.ReceiptReadMenuResponse, error) {
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
	if sendDto.MenuStatus != 0 {
		setUpdateSet = append(setUpdateSet, bson.E{Key: "menuStatus", Value: sendDto.MenuStatus})
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

	return changeMenuEntity2ReceiptReadDto(*saveItem), nil
} 



/////////////////////////
//	  Delete
/////////////////////////

//Model 로직은 구현했지만 실제 삭제는 사용하지 않기에 모델에서만 구현체를 둔다.


/////////////////////////
//	  Utils
/////////////////////////

//메뉴 엔티티를 읽기 DTO로 변환시키는 함수
func changeMenuEntity2ReceiptReadDto(entity receipt_model.MenuEntity) *dto.ReceiptReadMenuResponse {
	var subMenu []dto.SubMenuRequest
	for _, val := range entity.SubMenu {
		item := dto.SubMenuRequest{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	return &dto.ReceiptReadMenuResponse{Id: entity.Id, Name:  entity.Name, MenuStatus: entity.MenuStatus, Price: entity.Price, Event: entity.Event, MenuCategory: entity.MenuCategory, SubMenu:  subMenu, FoodEtcInfo: dto.FoodEtcInfoRequest(entity.FoodEtcInfo), CreateDate: entity.CreateDate, UpdateDate: entity.UpdateDate}
}

//공통 MenuRequest 데이터 중 일부를 Entity랑 매핑하는 함수
func changeMenuRequest2Entity(sendDto dto.MenuRequest) receipt_model.MenuEntity {
	var menuEntity receipt_model.MenuEntity
	
	menuEntity.Name  = sendDto.Name
	menuEntity.MenuStatus = sendDto.MenuStatus
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

