package rating_controller

import (
	"net/http"
	"wemade_project/dto"
	rating_service "wemade_project/service/rating"

	"github.com/gin-gonic/gin"
)

/////////////////////////
//	  Struct
/////////////////////////

type RatingController struct {
	ratingService rating_service.RatingService
}


/////////////////////////
//	  Init func
/////////////////////////

//생성자 역할 함수
func InitWithSelf(ratingService rating_service.RatingService) RatingController {
	return RatingController{ratingService: ratingService}
}



/////////////////////////
//	  Add (Create)
/////////////////////////

//메뉴 등록 함수
func (rc *RatingController) AddRating() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		//요청 데이터를 확인한다.
		var addReq dto.CreateRatingRequest
		
		if err := ginCtx.ShouldBindJSON(&addReq); err != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: err.Error()}
			ginCtx.JSON(http.StatusBadRequest, errorBody )	

			//예외 바인딩
			// handleBindError(ginCtx, addMenuReq, "checkerMenuEvent", err)
			return 
		}

		//전달된 아이템 등록 처리
		result, mongoErr := rc.ratingService.AddRatingItem(addReq)

		//조회에 에러가 난 경우
		if mongoErr != nil {
			errorBody := dto.ResponseBody{Result: false, Msg: mongoErr.Error()}
			ginCtx.JSON(http.StatusBadGateway, errorBody )	
			return
		}
	
		ginCtx.JSON(http.StatusOK, dto.ResponseBody{Result: true, Data: result})
	}
}