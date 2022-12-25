package order_list_controller

import (
	"net/http"
	"wemade_project/dto"
	order_list_service "wemade_project/service/order_list"

	"github.com/gin-gonic/gin"
)

/////////////////////////
//	  Struct
/////////////////////////

type OrderListController struct {
	orderListService order_list_service.OrderListService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(orderListService order_list_service.OrderListService) OrderListController {
	return OrderListController{orderListService: orderListService}
}


/////////////////////////
//	  Add (Create)
/////////////////////////

//주문 접수 
func (c *OrderListController) AddOrderListItem() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addReq dto.CreateOrderListRequest

		if err := ginCtx.ShouldBindJSON(&addReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	
		}

		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.AddItem(addReq)

		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}
	
		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}


/////////////////////////
//	  Update
/////////////////////////

func (c *OrderListController) UpdateOrderList()  gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var updateReq dto.UpdateOrderListRequest

		if err := ginCtx.ShouldBindJSON(&updateReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	
		}

		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.UpdateOrderList(updateReq)
		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}


		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}