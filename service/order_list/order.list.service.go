package order_list_service

import (
	"errors"
	"time"
	"wemade_project/dto"
	order_list_enums "wemade_project/enums/order"
	order_list_model "wemade_project/model/order"
	menu_service "wemade_project/service/menu"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

/////////////////////////
//	  Struct
/////////////////////////

type OrderListService struct {
	orderListCollection order_list_model.OrderListCollection
	menuService menu_service.MenuService
}

/////////////////////////
//	  Init 
/////////////////////////


func InitWithSelf(model order_list_model.OrderListCollection, menuService menu_service.MenuService) OrderListService {
	return OrderListService{orderListCollection: model, menuService:  menuService }
}


/////////////////////////
//	  Find
/////////////////////////

/**
* Order User Id로 데이터를 조회하는 함수
*/
func (s *OrderListService) Find4OrderUserId(orderUserId string, sortOpt string) ([]*dto.NomalReadOrderListResponse, error) {
	//등록된 아이템을 반환하기 위해 조회
	findItem, err1 := s.orderListCollection.Find4OrderUserIdAndStatus(orderUserId, sortOpt)
	if err1 != nil {
		return nil, err1
	}

	var itemList []*dto.NomalReadOrderListResponse
	for _, item := range findItem {
		itemList = append(itemList, s.changeEntity2FullReadDto(*item))
	}

	return itemList, nil
}


func (s *OrderListService) Find4All() ([]*dto.NomalReadOrderListResponse, error)  {
	//등록된 아이템을 반환하기 위해 조회
	findItem, err1 := s.orderListCollection.FindByAll()
	if err1 != nil {
		return nil, err1
	}

	var itemList []*dto.NomalReadOrderListResponse
	for _, item := range findItem {
		itemList = append(itemList, s.changeEntity2FullReadDto(*item))
	}

	return itemList, nil
}

/////////////////////////
//	  Add Data
/////////////////////////

//Add Order item 
/**
* 주문 요청 등록
*/
func (s *OrderListService) AddItem(addDto dto.CreateOrderListRequest) (*dto.NomalReadOrderListResponse, error) {
	//Model 조립
	entity := changeCreateOrderListDto2Entity(addDto)

	//데이터 등록 처리
	result, err := s.orderListCollection.AddEntity(entity);
	if err != nil {
		return nil, err
	}

	//등록된 아이템을 반환하기 위해 조회
	saveItem, err1 := s.orderListCollection.FindByObjectId(result.InsertedID); 
	if err1 != nil {
		return nil, err1
	}

	return s.changeEntity2NormalReadDto(*saveItem), nil
}

/////////////////////////
//	  Update 
/////////////////////////

/**
* 주문 접수 후 메뉴 수정 처리
*/
func (s *OrderListService) UpdateOrderList4Menu(sendDto dto.UpdateOrderList4MenuRequest) (*dto.NomalReadOrderListResponse, error) {
	//Id를 통해 수정 대상 메뉴 아이템을 찾아온다.
	findMenu, findErr := s.orderListCollection.FindByOrderId(sendDto.OrderId)
	if findErr != nil {
		return nil, findErr
	}

	//주문 접수 상태가 아니라면 변경 실패 처리한다.
	if (findMenu.OrderStatus != order_list_enums.OrderReceipt) {
		return nil, errors.New("Order change fail.")
	}
	
	//먼저 주문 상태를 변경하고 업데이트 처리한다.
	var setUpdateSet bson.D
	setUpdateSet = append(setUpdateSet, bson.E{Key: "orderStatus", Value: order_list_enums.OrderAddChange})
	setUpdateSet = append(setUpdateSet, bson.E{Key: "updateDate", Value: time.Now()})
	s.orderListCollection.UpdateEntity(findMenu.ID, setUpdateSet)

	//신규 주문 접수 상태를 만들어준다.
	addDto :=dto.CreateOrderListRequest{OrderUserId: findMenu.OrderUserId, OrderMenuList: sendDto.OrderMenu}
	return s.AddItem(addDto)
}


//상태 업데이트
func (s *OrderListService) UpdateOrderList4Status(sendDto dto.UpdateOrderList4StatusRequest) (*dto.NomalReadOrderListResponse, error) {
	//Id를 통해 수정 대상 메뉴 아이템을 찾아온다.
	findMenu, findErr := s.orderListCollection.FindByOrderId(sendDto.OrderId)
	if findErr != nil {
		return nil, findErr
	}

	//먼저 주문 상태를 변경하고 업데이트 처리한다.
	var setUpdateSet bson.D
	setUpdateSet = append(setUpdateSet, bson.E{Key: "orderStatus", Value: sendDto.OrderStatus})
	setUpdateSet = append(setUpdateSet, bson.E{Key: "updateDate", Value: time.Now()})
	s.orderListCollection.UpdateEntity(findMenu.ID, setUpdateSet)


	//등록된 아이템을 반환하기 위해 조회
	saveItem, err1 := s.orderListCollection.FindByObjectId(findMenu.ID); 
	if err1 != nil {
		return nil, err1
	}
	
	return s.changeEntity2NormalReadDto(*saveItem), nil
}


/////////////////////////
//	  Utils
/////////////////////////

//Creat Request dto 2 Entitiy Mapping
func changeCreateOrderListDto2Entity(sendDto dto.CreateOrderListRequest) order_list_model.OrderListEntity {
	var entity order_list_model.OrderListEntity

	entity.OrderId = uuid.Must(uuid.NewV4()).String()
	entity.OrderUserId = sendDto.OrderUserId
	entity.OrderMenu = sendDto.OrderMenuList
	entity.OrderStatus = order_list_enums.OrderReceipt //초기 상태는 주문 요청으로 기본 접수한다.
	entity.CreateDate = time.Now()
	entity.UpdateDate = entity.CreateDate

	return entity
}


//Entity 2 Normal read dto Mapping
func (s *OrderListService) changeEntity2NormalReadDto(entity order_list_model.OrderListEntity) *dto.NomalReadOrderListResponse {
	return &dto.NomalReadOrderListResponse{OrderId: entity.OrderId, OrderUserId: entity.OrderUserId, OrderStatus:  entity.OrderStatus, TotalPrice: s.GetTotalPrice(entity.OrderMenu)}
}

func (s *OrderListService) changeEntity2FullReadDto(entity order_list_model.OrderListEntity) *dto.NomalReadOrderListResponse {

	var menuList []dto.NormalReadMenuResponse
	for _, item := range entity.OrderMenu {
		menuData, _ := s.menuService.Find4MenuId(item)
		menuList = append(menuList, dto.NormalReadMenuResponse{Name: menuData.Name, Price: menuData.Price, MenuCategory: menuData.MenuCategory, FoodEtcInfo: menuData.FoodEtcInfo} )
	}
	return &dto.NomalReadOrderListResponse{OrderId: entity.OrderId, OrderUserId: entity.OrderUserId, OrderStatus:  entity.OrderStatus, TotalPrice: s.GetTotalPrice(entity.OrderMenu), OrderMenu:menuList  }
}

 
//전달된 주문 리스트의 오더 리스트에서 총 가격을 구해오는 함수
func (s *OrderListService) GetTotalPrice(menuIdList []string) int {
	var totalPrice int
	for _, menuId := range menuIdList {
		menuEntity, _ := s.menuService.Find4MenuId(menuId)
		totalPrice +=  menuEntity.Price
	}
	return totalPrice
}



