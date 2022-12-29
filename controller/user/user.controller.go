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

// @Summary 사용자 회원 가입
// @Description 사용자 서비스 가입 API
// @Tags Account-User-API
// @Success 200 {object} dto.NomalReadUserResponse
// @Accept  json
// @Produce  json
// @Param dto body dto.CreateUserRequest true "일반 사용자 등록용 DTO. dto.CreateUserRequest 객체 참고"
// @Router /api/v1/account/user/add [post]
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