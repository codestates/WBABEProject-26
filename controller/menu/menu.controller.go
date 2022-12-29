package menu_controller

import (
	"fmt"
	"net/http"

	"wemade_project/dto"
	menu_dto "wemade_project/dto"
	menu_enums "wemade_project/enums/menu"
	menu_service "wemade_project/service/menu"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuController struct {
	menuService menu_service.MenuService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(menuService menu_service.MenuService) MenuController {
	return MenuController{menuService: menuService}
}


/////////////////////////
//	  Get (Read)
/////////////////////////


// @Summary 메뉴 조회
// @Description Menu id를 전달하면 메뉴의 상세 정보 및 해당 메뉴의 평점 및 리뷰 정보를 제공하는 API
// @Tags Menu-API
// @Success 200 {object} dto.ReadMenuRatingResponse
// @Accept  json
// @Produce  json
// @Param menu_id path string true "Menu id <= 메뉴 고유 id값 "
// @Router /api/v1/store/menu/get/{menu_id} [get]
func (mc *MenuController) GetMenu4MenuId() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		menuId := ginCtx.Param("menu_id")

		result, mongoErr := mc.menuService.Find4MenuId(menuId)
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}

		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}


// @Summary 메뉴 리스트 조회
// @Description 메뉴 리스트 조회를 하는 API
// @Tags Menu-API
// @Success 200 {array} dto.HalfReadMenuResponse
// @Accept  json
// @Produce  json
// @Router /api/v1/store/menu/get/ [get]
func (mc *MenuController) GetMenuList() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		result, mongoErr := mc.menuService.FindMenuList()
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


// @Summary 메뉴 등록
// @Description 메뉴를 등록하는 함수
// @Tags Menu-API
// @Success 200 {object} dto.HalfReadMenuResponse
// @Accept  json
// @Produce  json
// @Param dto body menu_dto.CreateMenuRequest true "메뉴 등록용 DTO. dto.CreateMenuRequest 객체 참고"
// @Router /api/v1/store/menu/add [post]
func (mc *MenuController) AddMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addMenuReq menu_dto.CreateMenuRequest;

		if err := ginCtx.ShouldBindJSON(&addMenuReq); err != nil {
			fmt.Println("err = ", err )
			//예외 바인딩
			handleBindError(ginCtx, addMenuReq, "checkerMenuEvent", err)
			return 
		}

		//전달된 아이템 등록 처리
		result, mongoErr := mc.menuService.AddMenuItem(addMenuReq)

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
//	  Update (Update)
/////////////////////////

// @Summary 메뉴 수정
// @Description 메뉴를 수정하는 함수
// @Tags Menu-API
// @Success 200 {object} dto.UpdateMenuRequest
// @Accept  json
// @Produce  json
// @Param dto body menu_dto.UpdateMenuRequest true "메뉴 수정용 DTO. dto.UpdateMenuRequest 객체 참고"
// @Router /api/v1/store/menu/delete [put]
func (mc *MenuController) UpdateMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//UpdateMenuRequest
		var updateMenuReq menu_dto.UpdateMenuRequest

		if err := ginCtx.BindJSON(&updateMenuReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err}
			ginCtx.JSON(http.StatusBadGateway, errorBody )
			return
		}

		result, mongoErr := mc.menuService.UpdateMenuItem(updateMenuReq);
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}

		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}


/////////////////////////
//	  Delete (Delete)
/////////////////////////

// @Summary 메뉴 삭제
// @Description Menu 아이템을 논리적으로 삭제하는 함수 (물리적 X)
// @Tags Menu-API
// @Success 200 {object} dto.ResponseBody
// @Accept  json
// @Produce  json
// @Param menu_id path string true "Menu id <= 메뉴 고유 id값 "
// @Router /api/v1/store/menu/delete/{menu_id} [delete]
func (mc *MenuController) DeleteMenu4Logical() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		updateDto := dto.UpdateMenuRequest{Id: ginCtx.Param("menu_id"), MenuStatus: menu_enums.MSS_Delete }

		_, mongoErr := mc.menuService.UpdateMenuItem(updateDto);
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}

		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: "Menu delete success."})
	}
}


/////////////////////////
//	  Error Handler
/////////////////////////


func handleBindError(c *gin.Context, obj interface{}, tag string, err error) {
	var errs []gin.H
	switch err.(type) {
		case validator.ValidationErrors:	
			vErrs := err.(validator.ValidationErrors)
			for _, vErr := range vErrs {
				var message string
				value := vErr.Value()
				switch vErr.ActualTag() {
					case "checkerMenuEvent":
						message = fmt.Sprintln("Menu Event Code is wrong. Check Value. ", value)
					default:
            			message = err.Error()    
				}
				errs = append(errs, gin.H{
                	"field":   vErr.Field(),
                	"value":   value,
                	"message": message,
            	})
			}
		default: 
			errs = append(errs, gin.H{
                	"message": err.Error(),
            })
	}

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"errors": errs,
	})	 
}

