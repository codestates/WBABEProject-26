package user_controller

import (
	"net/http"
	"wemade_project/dto"
	user_service "wemade_project/service/user"

	"github.com/gin-gonic/gin"
)

/////////////////////////
//	  Struct
/////////////////////////

type UserController struct {
	userService user_service.UserService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(userService user_service.UserService) UserController {
	return UserController{userService: userService}
}


/////////////////////////
//	  Add (Create)
/////////////////////////

//메뉴 등록 함수
func (uc *UserController) AddUser() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addUserReq dto.CreateUserRequest
		
		if err := ginCtx.ShouldBindJSON(&addUserReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	

			//예외 바인딩
			// handleBindError(ginCtx, addMenuReq, "checkerMenuEvent", err)
			return 
		}

		//전달된 아이템 등록 처리
		result, mongoErr := uc.userService.AddMenuItem(addUserReq)

		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}
	
		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}