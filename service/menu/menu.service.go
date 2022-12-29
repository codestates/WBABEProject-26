package menu_service

import (
	"time"

	"wemade_project/dto"
	menu_model "wemade_project/model/menu"
	rating_service "wemade_project/service/rating"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuService struct {
	memuCollection menu_model.MenuCollection
	ratingService rating_service.RatingService
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model menu_model.MenuCollection, ratingService rating_service.RatingService) MenuService {
	return MenuService{memuCollection: model, ratingService:  ratingService }
}


/////////////////////////
//	  Find Data
/////////////////////////

/**
* 메뉴 id(고유Id, _id 아님)로 엔티티 데이터를 조회하는 함수
*/
func (ms *MenuService) Find4MenuId(menuId string) (*dto.ReadMenuRatingResponse, error) {
	findMenu, menuFindErr := ms.memuCollection.FindEntity2MenuId(menuId)
	if menuFindErr != nil {
		return nil, menuFindErr
	}

	//메뉴 평점도 같이 가지고 온다.
	findRating, raringErr := ms.ratingService.FindList4MenuId(findMenu.Id)
	if raringErr != nil {
		return nil, raringErr
	}

	return changeEntity2MenuRatingDto(*findMenu, findRating), nil
}

/**
* 메뉴 리스트 조회
*/
func (ms *MenuService) FindMenuList() ([]dto.HalfReadMenuResponse, error) {
	menuList, findErr := ms.memuCollection.FindEntityList2All()
	if findErr != nil {
		return nil, findErr
	}

	var result []dto.HalfReadMenuResponse
	for _, menuItem := range menuList {
		result = append(result, *ms.ChangeEntity2HalfReadDto(*menuItem))
	}
	
	return result, nil
}

/////////////////////////
//	  Add Data
/////////////////////////

func (ms *MenuService) AddMenuItem(addDto dto.CreateMenuRequest) (*dto.HalfReadMenuResponse, error) {
	//Model 조립
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

	return ms.ChangeEntity2HalfReadDto(*saveItem), nil
}


/////////////////////////
//	  Update
/////////////////////////

//메뉴 업데이트
func (ms *MenuService) UpdateMenuItem(sendDto dto.UpdateMenuRequest) (*dto.HalfReadMenuResponse, error) {
	//Id를 통해 수정 대상 메뉴 아이템을 찾아온다.
	findMenu, findErr := ms.memuCollection.FindEntity2MenuId(sendDto.Id)
	if findErr != nil {
		return nil, findErr
	}
	
	//수정할 데이터 셋
	var setUpdateSet bson.D
	/*
	모든 필드에 대해서 기본값이 아님을 체크하는 방법은 안티 패턴인 것 같습니다.
	메뉴의 필드가 계속해서 늘어난다면, 중복된 코드가 엄청나게 늘어날 것입니다.
	*/
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

	return ms.ChangeEntity2HalfReadDto(*saveItem), nil
} 



/////////////////////////
//	  Delete
/////////////////////////

//Model 로직은 구현했지만 실제 삭제는 사용하지 않기에 모델에서만 구현체를 둔다.


/////////////////////////
//	  Utils
/////////////////////////

//메뉴 엔티티를 읽기 DTO로 변환시키는 함수
func (ms *MenuService) ChangeEntity2HalfReadDto(entity menu_model.MenuEntity) *dto.HalfReadMenuResponse {
	var subMenu []dto.SubMenuRequest
	for _, val := range entity.SubMenu {
		item := dto.SubMenuRequest{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	/*
	return 부분의 가독성이 매우 떨어지는 것 같습니다. 따로 변수를 선언에 할당하거나, 적절한 줄바꿈이 필요해보입니다.
	*/
	return &dto.HalfReadMenuResponse{Id: entity.Id, Name:  entity.Name, MenuStatus: entity.MenuStatus, Price: entity.Price, Event: entity.Event, MenuCategory: entity.MenuCategory, SubMenu:  subMenu, FoodEtcInfo: dto.FoodEtcInfoRequest(entity.FoodEtcInfo), CreateDate: entity.CreateDate, UpdateDate: entity.UpdateDate}
}

func changeEntity2MenuRatingDto(entity menu_model.MenuEntity, rating []*dto.FullReadRatingResponse) *dto.ReadMenuRatingResponse {
	var subMenu []dto.SubMenuRequest
	for _, val := range entity.SubMenu {
		item := dto.SubMenuRequest{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}

	//(평점 부분) Dto를 위한 형 변환 작업 처리
	var tmpData []dto.FullReadRatingResponse
	for _, item := range rating {
		tmpData = append(tmpData, *(item))
	}
	
	return &dto.ReadMenuRatingResponse{Id: entity.Id, Name:  entity.Name, MenuStatus: entity.MenuStatus, Price: entity.Price, Event: entity.Event, MenuCategory: entity.MenuCategory, 
		SubMenu:  subMenu, FoodEtcInfo: dto.FoodEtcInfoRequest(entity.FoodEtcInfo), Rating: tmpData, CreateDate: entity.CreateDate, UpdateDate: entity.UpdateDate}
}



//공통 MenuRequest 데이터 중 일부를 Entity랑 매핑하는 함수
func changeMenuRequest2Entity(sendDto dto.MenuRequest) menu_model.MenuEntity {
	var menuEntity menu_model.MenuEntity
	
	menuEntity.Name  = sendDto.Name
	menuEntity.MenuStatus = sendDto.MenuStatus
	menuEntity.Price = sendDto.Price
	menuEntity.Event = sendDto.Event
	menuEntity.MenuCategory = sendDto.MenuCategory

	menuEntity.SubMenu = changeSubMenu4Dto2Entity(sendDto.SubMenu)
	menuEntity.FoodEtcInfo = menu_model.FoodEtcInfo(sendDto.FoodEtcInfo)

	return menuEntity;
}


//Menu 안의 SubMenu 의 dto형태를 entity 형태로 맞춰주는 함수
func changeSubMenu4Dto2Entity(sendDto []dto.SubMenuRequest)[]menu_model.SubMenu {
	var subMenu []menu_model.SubMenu
	for _, val := range sendDto {
		item := menu_model.SubMenu{SubMenuName: val.SubMenuName, Name: val.Name, Price: val.Price}
		subMenu = append(subMenu, item)
	}
	return subMenu
}

