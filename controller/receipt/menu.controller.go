package receipt_controller

import (
	"fmt"
	"net/http"

	"wemade_project/dto"
	receipt_dto "wemade_project/dto"
	receipt_enums "wemade_project/enums/receipt"
	receipt_service "wemade_project/service/receipt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

/////////////////////////
//	  Struct
/////////////////////////

type MenuController struct {
	menuService receipt_service.MenuService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(menuService receipt_service.MenuService) MenuController {
	return MenuController{menuService: menuService}
}


/////////////////////////
//	  Get (Read)
/////////////////////////

//메뉴 조회
func (mc *MenuController) GetMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ginCtx.JSON(http.StatusOK, gin.H{"Message" : "Get Menu"})
	}
}

/////////////////////////
//	  Add (Create)
/////////////////////////

//메뉴 등록 함수
func (mc *MenuController) AddMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addMenuReq receipt_dto.CreateMenuRequest;

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

//메뉴 수정
func (mc *MenuController) UpdateMenu() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//UpdateMenuRequest
		var updateMenuReq receipt_dto.UpdateMenuRequest

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

//Menu 아이템을 논리적으로 삭제하는 함수
func (mc *MenuController) DeleteMenu4Logical() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		updateDto := dto.UpdateMenuRequest{Id: ginCtx.Param("menu_id"), MenuStatus: receipt_enums.MSS_Delete }

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

/////Error bind
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

