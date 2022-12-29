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


// @Summary 평점 및 리뷰 등록
// @Description 소비자가 리뷰 및 평점을 등록하는 API
// @Tags Rating-API
// @Success 200 {object} dto.FullReadRatingResponse
// @Accept  json
// @Produce  json
// @Param dto body dto.CreateRatingRequest true "리뷰 평점 등록용 DTO. dto.CreateRatingRequest 객체 참고"
// @Router /api/v1/rating/add [post]
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