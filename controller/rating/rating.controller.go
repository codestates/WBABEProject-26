package rating_controller

import rating_service "wemade_project/service/rating"

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
