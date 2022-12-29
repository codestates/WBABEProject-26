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
//	  Find
/////////////////////////


// @Summary 주문 내역 조회 (사용자)
// @Description User id를 전달하면 해당 사용자의 주문 내역을 제공하는 API
// @Tags OrderList-Order-API
// @Success 200 {array} dto.NomalReadOrderListResponse
// @Accept  json
// @Produce  json
// @Param user_id path string true "User id => User entity ID(_id) 값 "
// @Router /api/v1/order_list/order/user/{user_id} [get]
func (c *OrderListController) Find4OrderUserId() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		userId := ginCtx.Param("user_id")
		sortOption := ginCtx.DefaultQuery("on_stage", "no")

		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.Find4OrderUserId(userId, sortOption)

		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}
	
		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}

// @Summary 주문 내역 조회
// @Description 주문 내역 전체를 조회하는 API
// @Tags OrderList-Order-API
// @Success 200 {array} dto.NomalReadOrderListResponse
// @Accept  json
// @Produce  json
// @Router /api/v1/order_list/order [get]
func (c *OrderListController) Find4All() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		
		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.Find4All()

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
//	  Add (Create)
/////////////////////////


// @Summary 주문 접수
// @Description 주문을 접수하는 API
// @Tags OrderList-Order-API
// @Success 200 {object} dto.NomalReadOrderListResponse
// @Accept  json
// @Produce  json
// @Param dto body dto.CreateOrderListRequest true "주문 접수용 DTO. dto.CreateOrderListRequest 객체 참고"
// @Router /api/v1/order_list/order/add [post]
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

// @Summary 주문 내역 갱신 (소비자 입장)
// @Description 소비자가 메뉴 추가 등 주문을 갱신하는 API
// @Tags OrderList-Order-API
// @Success 200 {object} dto.NomalReadOrderListResponse
// @Accept  json
// @Produce  json
// @Param dto body dto.UpdateOrderList4MenuRequest true "주문 내용 갱신용 DTO. dto.UpdateOrderList4MenuRequest 객체 참고"
// @Router /api/v1/order_list/order/menu/update [put]
func (c *OrderListController) UpdateOrderList4Menu()  gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var updateReq dto.UpdateOrderList4MenuRequest

		if err := ginCtx.ShouldBindJSON(&updateReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	
		}

		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.UpdateOrderList4Menu(updateReq)
		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}


		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}


// @Summary 주문 내역 갱신 (업주 입장)
// @Description 업주가 주문 상태 정보 등을 갱신하는 API
// @Tags OrderList-Order-API
// @Success 200 {object} dto.UpdateOrderList4StatusRequest
// @Accept  json
// @Produce  json
// @Param dto body dto.UpdateOrderList4StatusRequest true "주문 상태 갱신용 DTO. dto.UpdateOrderList4StatusRequest 객체 참고"
// @Router /api/v1/order_list/order/status/update [put]
func (c *OrderListController) UpdateOrderList4Status()  gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var updateReq dto.UpdateOrderList4StatusRequest

		if err := ginCtx.ShouldBindJSON(&updateReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	
		}

		//전달된 아이템 등록 처리
		result, mongoErr := c.orderListService.UpdateOrderList4Status(updateReq)
		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}


		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}