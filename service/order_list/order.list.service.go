package order_list_service

import (
	"time"
	"wemade_project/dto"
	order_list_enums "wemade_project/enums/order"
	order_list_model "wemade_project/model/order"
	menu_service "wemade_project/service/menu"

	"github.com/gofrs/uuid"
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

	return changeEntity2NormalReadDto(*saveItem), nil
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
func changeEntity2NormalReadDto(entity order_list_model.OrderListEntity) *dto.NomalReadOrderListResponse {
	return &dto.NomalReadOrderListResponse{OrderId: entity.OrderId, OrderUserId: entity.OrderUserId, OrderStatus:  entity.OrderStatus, TotalPrice: 0 }
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



